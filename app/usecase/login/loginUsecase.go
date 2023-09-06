package login

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Danangoffic/configib/app/Tools"
	"github.com/Danangoffic/configib/app/models"
	"github.com/Danangoffic/configib/app/repositories/account"
	"github.com/Danangoffic/configib/app/repositories/devicelockdown"
	"github.com/Danangoffic/configib/app/repositories/lookup"
	"github.com/Danangoffic/configib/app/repositories/systemparameter"
	"github.com/Danangoffic/configib/app/repositories/user"
	"github.com/gin-gonic/gin"
)

type loginUsecase struct {
	accountRepo        account.Repository
	lookupRepo         lookup.Repository
	systemparamRepo    systemparameter.Repository
	userRepo           user.Repository
	devicelockdownRepo devicelockdown.Repository
}

func NewLoginUsecase(
	accountRepo account.Repository,
	lookupRepo lookup.Repository,
	systemparamRepo systemparameter.Repository,
	userRepo user.Repository,
	devicelockdownRepo devicelockdown.Repository,
) Usecase {
	return &loginUsecase{
		accountRepo:       accountRepo,
		lookupRepo:        lookupRepo,
		systemparamRepo:   systemparamRepo,
		userRepo:          userRepo,
		devicelocdownRepo: devicelockdownRepo,
	}
}

func (p *loginUsecase) DoLoginScrumV2(c *gin.Context, req models.LoginRequestScrum) {
	var billpayMethodType, transferMethodType, doneBy string
	billpayMethodType = req.BillpayMethodType.(string)

	transferMethodType = req.TransferMethodType.(string)

	lang := req.Lang.(string)
	sessionCode := req.SessionCode.(string)
	loginMobileNumber := req.LoginMobileNumber.(string)
	username := req.Username.(string)

	accType := req.AccType.(string)

	if isNum(username) {
		loginMobileNumber = username
	}
	// var loginPassword string
	loginPassword, ok := req.LoginPassword.(string)
	if !ok {
		log.Println("loginPassword is not of type string")
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "22",
			"responseMessage": "Login failed, please try again later",
		})
		return
	}
	easyPin := req.EasyPin.(string)
	ciamPassword := req.CiamPassword.(string)
	ciamPin := req.CiamPin.(string)
	deviceParam := req.DeviceParam.(string)
	activeOtp := req.ActiveOtp.(string)
	accessToken := req.AccessToken.(string)
	tokenId := req.TokenId.(string)

	additionalMap := make(map[string]interface{})
	additionalMap["accessToken"] = accessToken
	additionalMap["tokenId"] = tokenId

	challenge := req.E2EE_RANDOM.(string)

	// // toggle disable silent login
	toggleDisableSilentLogin, err := p.systemparamRepo.GetOnTheFly("TOOGLE", "DISABLED_SILENT_LOGIN")
	if err != nil {
		log.Fatalln("failed to get disabled silent login with ", err)
	}
	if strings.EqualFold(toggleDisableSilentLogin.Svalue, "YES") && req.EasyPin == nil && req.BillpayMethodType == nil && req.TransferMethodType == nil {
		log.Println("masuk block login finger print disabled")
		returnMap := map[string]string{}
		returnMap["responseCode"] = "01"
		returnMap["responseMessage"] = "Login Error"
		c.JSON(http.StatusBadRequest, returnMap)
		return
	}

	versionScope := req.VersionScope.(string)
	log.Println("versionScope : ", versionScope)

	deviceCodeParam := req.DeviceCodeParam.(string)
	tokenClient := req.TokenClient.(string)
	tokenServer := req.TokenServer.(string)
	ciamAPK := req.CiamAPK.(string)

	OBMParameterMap := Tools.UserCredentialsToOBM(req)
	log.Println("OBMParameterMap : ", OBMParameterMap)

	var smsPriority bool
	if req.SmsPriority == nil {
		smsPriority = false
	} else {
		smsPriority = req.SmsPriority.(bool)
	}
	log.Println("has sms priority : ", smsPriority)

	if Tools.NewAppFactory().IsEmpty(challenge) && Tools.NewAppFactory().IsEmpty(sessionCode) {
		c.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "01",
			"responseMessage": "failed",
		})
		return
	}

	// var deviceCodeMap map[string]string
	deviceCodeMap := make(map[string]string)
	defer func() {
		if err := recover(); err != nil {
			log.Println(err, err)
		}
	}()

	if tokenServer != "" && len(tokenServer) > 0 {
		tokenClient, err := models.OracleCodec(tokenClient, tokenServer)
		if err != nil {
			log.Println(err)
		}
		tokenServer, err = models.OracleCodec(tokenServer, tokenClient)
		if err != nil {
			log.Println(err)
		}
		// tokenClient = esapi.Encoder().EncodeForSQL(oracleCodec, tokenClient)
		// tokenServer = esapi.Encoder().EncodeForSQL(oracleCodec, tokenServer)
		log.Println("[" + sessionCode + "] - tokenClient : " + tokenClient)
		log.Println("[" + sessionCode + "] - tokenServer : " + tokenServer)

		deviceCodeMap = Tools.NewAppFactory().GetDeviceLockToken(tokenClient, tokenServer)
		deviceCodeParam = deviceCodeMap["hash"]
	}

	if sessionCode == "" {
		// Generate a new session code.
		sessionCode = strconv.FormatInt(time.Now().Unix(), 10)

		// 	// Log the session code.
		log.Println("sessionCode: ", sessionCode)
	}

	var (
		methodLogin string
	)
	isSSLApps, err := p.systemparamRepo.GetOnTheFly("FLAG", "IS_SSL_APPS")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("isSSLApps : ", isSSLApps.Svalue)

	var (
		uMobileNumber    string = ""
		uMobileNumberSSL string = ""
	)

	if strings.EqualFold(isSSLApps.Svalue, "true" && req.ActiveOtp != nil) {
		uMobileNumberSSL = loginMobileNumber
	} else {
		uMobileNumber = loginMobileNumber
	}
	log.Println("mobile number ssl : ", uMobileNumberSSL)
	log.Println("mobile number : ", uMobileNumber)

	var (
		password                   string = loginPassword
		passcode                   string = ""
		mobileNumber               string = ""
		isMobile                   bool
		usrChech                   models.User
		encryptedMobileNumber      string
		result                     map[string]interface{}
		customer                   map[string]interface{}
		isCheckingCore             bool = false
		usernameEmptyMobile        string
		transactionReferenceNumber string = ""
	)

	if strings.EqualFold(isSSLApps.Svalue, "true") {
		transactionReferenceNumber = Tools.FileHiloGenerator().GenerateUniqueCode("SL")
	} else {
		transactionReferenceNumber = Tools.FileHiloGenerator().GenerateUniqueCode("MB")
	}
	log.Println("transactionReferenceNumber : ", transactionReferenceNumber)

	if len(deviceCodeParam) > 0 {
		log.Printf("[%s] - LOGIN VIA LOCKDOWN DEVICE", sessionCode)

		deviceCodeParam, _ = models.OracleCodec(deviceCodeParam, deviceParam)
		log.Printf("[%s] - deviceCodeParam : %s", sessionCode, deviceCodeParam)

		deviceLockdownObj, err := p.devicelockdownRepo.FindByDeviceCode(deviceCodeParam)
		if err != nil {
			log.Println("failed get device lockdown with : ", err)
		}
		log.Println("deviceLockdownObject : ", deviceLockdownObj)
		if deviceLockdownObj.ID == 0 {
			log.Println("device lockdown not found")
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode":    "06",
				"responseMessage": "device lockdown not found",
			})
			return
		}

		var isBypassPassword bool = false

		encryptedMobileNumber = deviceLockdownObj.AppUser.MobileNumber.(string)

		userTemp, err := p.userRepo.FindByID(deviceLockdownObj.UserID)
		if err != nil {
			log.Println("failed to find user by device lockdown data obj with : ", err)
		}

		if checkDormant(userTemp) {
			log.Printf("[%s] - loginUser is dormant!", sessionCode)
			var returnMap gin.H
			returnMap["responseCode"] = "05"
			returnMap["responseMessage"] = "user is dormant"
			if userTemp.MobileNumber.(string) != "" && userTemp.MobileNumber != nil {
				returnMap["isMobileNumberRegistered"] = true
			} else {
				returnMap["isMobileNumberRegistered"] = false
			}
			c.JSON(http.StatusBadRequest, returnMap)
		}

		// close face recognition caused lib error
		// var (
		// 	faceRecognitionMap gin.H
		// 	faceRecognitionOri string = req.FaceRecognition.(string)
		// 	faceRecognition    string = faceRecognitionOri
		// )

		// if len(faceRecognition) > 0 || faceRecognition != "" {
		// 	var (
		// 		orientation          string = req.Orientation.(string)
		// 		orientationDegree, _        = strconv.Atoi(orientation)
		// 		flipImage            bool   = req.FlipImage.(bool)
		// 	)
		// 	if len(faceRecognitionOri) > 0 && faceRecognitionOri != "" && "undefined" != faceRecognitionOri {
		// 		regexStr := "(\n|\r)"
		// 		faceRecognition = strings.ReplaceAll(faceRecognitionOri, "", regexStr)
		// 		if orientationDegree > 0 {
		// 			if flipImage {
		// 				// TODO: flipimage from appFactory
		// 				// faceRecognition
		// 			}
		// 		}
		// 	}
		// }

		var isCorporateDisableLogin bool = true
		// result =

	}
}

func checkDormant(loginUser models.AppUser) bool {
	log.Println("checkdormant user : ", loginUser)
	if loginUser.ID != 0 && strings.EqualFold(loginUser.Status.Code, "dormant") {
		log.Println("checkdormant user loginame : ", loginUser.LoginName)
		log.Println("checkdormant user status : ", loginUser.Status.Code)
		return true
	}
	return false
}
