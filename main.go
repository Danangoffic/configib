package main

import (
	"encoding/json"
	"log"

	"github.com/Danangoffic/configib/app/handlers"
	"github.com/Danangoffic/configib/app/repositories/account"
	"github.com/Danangoffic/configib/app/repositories/customer"
	"github.com/Danangoffic/configib/app/repositories/devicelockdown"
	"github.com/Danangoffic/configib/app/repositories/lookup"
	"github.com/Danangoffic/configib/app/repositories/status"
	"github.com/Danangoffic/configib/app/repositories/systemparameter"
	"github.com/Danangoffic/configib/app/repositories/user"
	"github.com/Danangoffic/configib/app/usecase/login"
	"github.com/Danangoffic/configib/config"
	"github.com/gin-gonic/gin"
)

func main() {
	configDB, err := config.SetupDb()
	if err != nil {
		log.Fatal("failed to connect with db with : ", err)
	}

	// load all repositories first here
	statusRepo := status.GetStatusRepository(configDB)
	accountRepo := account.GetAccountRepository(configDB)
	customerRepo := customer.GetCustomerRepository(configDB, statusRepo, accountRepo)
	userRepo := user.GetUserRepository(configDB, statusRepo, customerRepo)
	systemParamRepo := systemparameter.GetSystemParamRepository(configDB)
	lookupRepo := lookup.GetLookupRepository(configDB)
	// accountRepo := account.GetAccountRepository(configDB)
	devicelockdownRepo := devicelockdown.GetDeviceLockdownRepository(configDB)

	// load all usecases here after repositories
	loginUsecase := login.NewLoginUsecase(accountRepo, lookupRepo, systemParamRepo, userRepo, devicelockdownRepo)

	// load handlers
	handlerMod := handlers.HandlerModel{LoginUsecase: loginUsecase}

	// log.Println("user status : ", user.Status)
	// log.Println("customer data : ", user.Customer)
	// log.Println("accounts : ", user.Customer.Accounts)

	r := gin.Default()
	v3Action := r.Group("/v3/action")
	v3Action.POST("/login-new", handlerMod.DoLoginScrumV2)

	r.GET("/ping", func(c *gin.Context) {
		user, err := userRepo.FindByMobileNumberIncludeForceReset("danangarif09")
		if err != nil {
			log.Println("failed to get user with : ", err)
		}
		var jsondata map[string]interface{}
		byteD, err := json.Marshal(user)
		json.Unmarshal(byteD, &jsondata)
		log.Println("user app : ", jsondata)
		c.JSON(200, jsondata)
		return
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
