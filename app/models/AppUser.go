package models

import (
	"encoding/json"
	"log"
	"strings"
)

type User struct {
	AppUser  AppUser
	Customer Customer
	Status   Status
}

type AppUser struct {
	// gorm.Model
	ID                 int64       `gorm:"column:ID" json:"id"`
	CustomerID         int64       `gorm:"column:CUSTOMER_ID" json:"customer_id"`
	Customer           Customer    `gorm:"foreignKey:CustomerID" json:"customer"`
	LoginName          string      `gorm:"column:LOGIN_NAME" json:"loginName"`
	LoginPassword      string      `gorm:"column:LOGIN_PASSWORD" json:"loginPassword"`
	PinCode            string      `gorm:"column:PIN_CODE" json:"pinCode"`
	MobileNumber       interface{} `gorm:"column:MOBILE_NUMBER" json:"mobileNumber"`
	UserLevel          int64       `gorm:"column:USER_LEVEL" json:"userLevel"`
	Title              string      `gorm:"column:TITLE" json:"title"`
	Name               string      `gorm:"column:NAME" json:"name"`
	RsaUserName        string      `gorm:"column:RSA_USER_NAME" json:"rsaUserName"`
	Email              string      `gorm:"column:EMAIL" json:"email"`
	StatusID           int64       `gorm:"column:STATUS" json:"status_id"`
	Status             Status      `gorm:"foreignKey:StatusID;references:ID" json:"status"`
	CreatedDate        interface{} `gorm:"column:CREATED_DATE" json:"createdDate"`
	CreatedBy          string      `gorm:"column:CREATED_BY" json:"createdBy"`
	LastUpdatedDate    interface{} `gorm:"column:LAST_UPDATED_DATE" json:"lastUpdatedDate"`
	LastUpdatedBy      string      `gorm:"column:LAST_UPDATED_BY" json:"lastUpdatedBy"`
	FirstLogin         string      `gorm:"column:FIRST_LOGIN" json:"firstLogin"`
	LastPasswordUpdate interface{} `gorm:"column:LAST_PASSWORD_UPDATE" json:"lastPasswordUpdate"`
	LastPinCodeUpdate  interface{} `gorm:"column:LAST_PIN_CODE_UPDATE" json:"lastPinCodeUpdate"`
	LastPassword       interface{} `gorm:"column:LAST_PASSWORD" json:"lastPassword"`
	LastPinCode        interface{} `gorm:"column:LAST_PIN_CODE" json:"lastPinCode"`
	PinCodeOBM         interface{} `gorm:"column:PIN_CODE_OBM" json:"pinCodeOBM"`
	PasswordOBM        interface{} `gorm:"column:PASSWORD_OBM" json:"passwordOBM"`
	EnablePrissa       string      `gorm:"column:ENABLE_PRISSA" json:"enablePrissa"`
	CustIDNuri         interface{} `gorm:"column:CUST_ID_NURI" json:"custIdNuri"`
	Language           interface{} `gorm:"column:LANG" json:"lang"`
	EmployeeNumber     interface{} `gorm:"column:EMPLOYEE_NUMBER" json:"employeeNumber"`
	LoginFailed        int         `gorm:"column:LOGIN_FAILED" json:"loginFailed"`
	LastActive         interface{} `gorm:"column:LAST_ACTIVE" json:"lastActive"`
	// CurrentSessionLoginHistory LoginHistory           `gorm:"references:UserID" json:"currentSessionLoginHistory"`
	// StatusMap    map[string]interface{} `json:"statusMap"`
	// ViaMobile    bool
	// CaptchaInput string
	// CaptchaID    string
	SessionCode interface{} `gorm:"column:SESSION_CODE" json:"sessionCode"`
	IPassport   interface{} `gorm:"column:IPASSPORT" json:"iPassport"`
	// UserToken    []map[string]interface{} `json:"userToken"`
	// UserPreference             Lookup                   `gorm:"foreignKey:UserPreference;references:ID" json:"userPreference"`
	// MobilNumberDecrypted      string
	// MobilNumberDecryptedPlan  string
	// MobileNumberMasking       string
	MaxRelease                int         `gorm:"column:MAX_RELEASE" json:"maxRelease"`
	LastReleasedDate          interface{} `gorm:"column:LAST_RELEASED_DATE" json:"lastReleasedDate"`
	TempData                  interface{} `gorm:"column:TEMP_DATA" json:"tempData"`
	UserApiKey                interface{} `gorm:"column:USER_API_KEY" json:"userApiKey"`
	FaceRecognition           interface{} `gorm:"column:FACE_RECOGNITION" json:"faceRecognition"`
	ReferralCode              interface{} `gorm:"column:REFERRAL_CODE" json:"referralCode"`
	UserApiKeySimobi          interface{} `gorm:"column:USER_API_KEY_SIMOBI" json:"userApiKeySimobi"`
	LastFaceRecognitionUpdate interface{} `gorm:"column:LAST_FACE_RECOGNITION_UPDATE" json:"lastFaceRecognitionUpdate"`
	ProfilePicture            interface{} `gorm:"column:PROFILE_PICTURE" json:"profilePicture"`
	IsFPrint                  interface{} `gorm:"column:ISFPRINT" json:"isFPrint"`
	IsFRecog                  interface{} `gorm:"column:ISFRECOG" json:"isFRecog"`
	IsAutoSave                interface{} `gorm:"column:ISAUTOSAVE" json:"isAutoSave"`
	AllowIB                   interface{} `gorm:"column:ALLOW_IB" json:"allowIB"`
	AutoSwitch                interface{} `gorm:"column:AUTO_SWITCH" json:"autoSwitch"`
	CiamUser                  interface{} `gorm:"column:CIAM_USER" json:"ciamUser"`
	// SecurityType               Lookup      `gorm:"foreignKey:SecurityType;references:ID" json:"securityType"`
}

type LoginRequestScrum struct {
	Username           interface{} `json:"username"`
	LoginPassword      interface{} `json:"loginPassword"`
	SessionCode        interface{} `json:"sessionCode"`
	LoginMobileNumber  interface{} `json:"loginMobileNumber"`
	AccType            interface{} `json:"accType"`
	Lang               interface{} `json:"lang"`
	EasyPin            interface{} `json:"easyPin"`
	CiamPassword       interface{} `json:"ciamPassword"`
	CiamPin            interface{} `json:"ciamPin"`
	DeviceParam        interface{} `json:"deviceParam"`
	ActiveOtp          interface{} `json:"activeOtp"`
	AccessToken        interface{} `json:"accessToken"`
	TokenId            interface{} `json:"tokenId"`
	FaceRecognition    interface{} `json:"faceRecognition"`
	Orientation        interface{} `json:"orientation"`
	FlipImage          interface{} `json:"flipImage"` // nanti parse ke bool
	AccountNumber      interface{} `json:"accountNumber"`
	BillpayMethodType  interface{} `json:"billpayMethodType"`
	TransferMethodType interface{} `json:"transferMethodType"`
	E2EE_RANDOM        interface{} `json:"E2EE_RANDOM"`
	VersionScope       interface{} `json:"versionScope"`
	DeviceCodeParam    interface{} `json:"deviceCodeParam"`
	TokenClient        interface{} `json:"tokenClient"`
	TokenServer        interface{} `json:"tokenServer"`
	CiamAPK            interface{} `json:"ciamAPK"`
	SmsPriority        interface{} `json:"smsPriority"`
}

func (u *AppUser) GetLanguageName() string {
	if strings.EqualFold(u.Language.(string), "en") {
		return "English"
	} else if strings.EqualFold(u.Language.(string), "id") {
		return "Indonesia"
	}
	return u.Language.(string)
}

// func (u *AppUser) GetStatusMapFromUserLogin() map[string]interface{} {
// 	var newStatusMap = make(map[string]interface{})
// 	if u.Status.ID > 0 {
// 		newStatusMap["id"] = u.Status.ID
// 		newStatusMap["type"] = u.Status.Type
// 		newStatusMap["code"] = u.Status.Code
// 		newStatusMap["name"] = u.Status.Name
// 		newStatusMap["description"] = u.Status.Description
// 	}
// 	return newStatusMap
// }

func (u *AppUser) GetUserToken() []string {
	var (
		userMobileNumber = u.MobileNumber.(string)
		// userSimasToken   = u.SecurityType
		userToken = []string{}
	)
	if userMobileNumber != "" && len(userMobileNumber) > 0 {
		userToken = append(userToken, "0")
	}
	// if userSimasToken.Code == "rsa" {
	// 	userToken = append(userToken, "1")
	// }
	return userToken
}

// func (u *AppUser) GetUserPreferenceMethod() string {
// 	var (
// 		userMobileNumber     = u.MobileNumber
// 		userPreferenceMethod = u.UserPreference
// 		userSimasToken       = u.SecurityType
// 		userPreference       string
// 	)
// 	if userPreferenceMethod.ID > 0 {
// 		userPreference = userPreferenceMethod.Code
// 	} else {
// 		if userMobileNumber != "" || userSimasToken.Code == "none" {
// 			userPreference = "0"
// 		} else if userSimasToken.Code == "rsa" && userMobileNumber == "" {
// 			userPreference = "1"
// 		} else if userMobileNumber != "" && userSimasToken.Code == "rsa" {
// 			userPreference = "2"
// 		}
// 	}
// 	return userPreference
// }

// func (u *AppUser) ToString() string {
// 	return fmt.Sprintln("User [id=", u.ID, ", loginName=", u.LoginName, ", status:", u.Status.Code, "\\n")
// }

func (u *AppUser) GetTempDataMap() map[string]interface{} {
	var tempDataMap = make(map[string]interface{})
	if u.TempData != "" {
		json.Unmarshal([]byte(u.TempData.(string)), &tempDataMap)
	}
	return tempDataMap
}
func (u *AppUser) SetTempDataMap(tempDataMap map[string]interface{}) {
	u.TempData = ""
	b, err := json.Marshal(tempDataMap)
	if err != nil {
		log.Println(err)
		return
	}
	u.TempData = string(b)
}
