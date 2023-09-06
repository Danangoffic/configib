package user

import (
	"errors"
	"log"

	"github.com/Danangoffic/configib/app/models"
	"github.com/Danangoffic/configib/app/repositories/customer"
	"github.com/Danangoffic/configib/app/repositories/status"
	"gorm.io/gorm"
)

type userRepo struct {
	DB         *gorm.DB
	statusRepo status.Repository
	custRepo   customer.Repository
}

func GetUserRepository(db *gorm.DB, statusrepo status.Repository, custRepo customer.Repository) Repository {
	return &userRepo{
		DB:         db,
		statusRepo: statusrepo,
		custRepo:   custRepo,
	}
}

func (r *userRepo) FindByID(id int64) (models.AppUser, error) {
	var appUser = models.AppUser{ID: id}
	err := r.DB.Preload("Customer").Unscoped().Find(&appUser).Error
	if err != nil {
		log.Println("failed to find app user with : ", err)
		return models.AppUser{}, err
	}
	if appUser.Name == "" {
		log.Println("user not found")
		return models.AppUser{}, errors.New("user not found")
	}
	if appUser.Status.ID == 0 {
		appUser.Status, _ = r.statusRepo.FindByID(appUser.StatusID)
	}
	if appUser.Customer.ID == 0 {
		appUser.Customer, _ = r.findCustomerByID(appUser.CustomerID)
	}

	return appUser, nil
}

func (r *userRepo) FindByMobileNumberIncludeForceReset(loginName string) (models.AppUser, error) {
	log.Println("loginName : ", loginName)

	var (
		appUser models.AppUser
	)

	err := r.DB.Table("APP_USER").Unscoped().Preload("Customer").
		Where("LOGIN_NAME = ?", loginName).
		Where(
			"STATUS IN (SELECT ID FROM STATUS S WHERE S.TYPE = 'user' AND (S.CODE = 'active' OR S.CODE = 'dormant' OR S.CODE = 'force_reset' OR S.CODE = 'inquiry_only'))").
		Find(&appUser).Error
	if err != nil {
		log.Println("failed to find a user by loginName with : ", err.Error())
		return models.AppUser{}, errors.New("user not found")
	}
	if appUser.ID == 0 {
		log.Println("user not found!")
		return models.AppUser{}, errors.New("user not found")
	}
	log.Println("user id : ", appUser.ID)
	if appUser.Status.ID == 0 {
		appUser.Status, _ = r.findStatusByID(appUser.StatusID)
	}
	if appUser.Customer.ID == 0 {
		appUser.Customer, _ = r.findCustomerByID(appUser.CustomerID)
	}
	return appUser, nil
}

func (r *userRepo) findCustomerByID(id int64) (models.Customer, error) {
	var customer = models.Customer{ID: id}
	log.Println("find customer id by : ", id)

	err := r.DB.Table("CUSTOMER").Unscoped().Find(&customer).Error
	if err != nil {
		log.Println("failed to find customer data with : ", err)
		return models.Customer{}, errors.New("failed to finc customer with " + err.Error())
	}

	if customer.ID == 0 {
		log.Println("customer not found")
		return models.Customer{}, errors.New("customer not found")
	}
	log.Println("cifcode : ", customer.CifCode)
	return customer, nil
}

func (r *userRepo) findStatusByID(id int64) (models.Status, error) {
	var status = models.Status{ID: id}

	err := r.DB.Table("STATUS").Unscoped().Find(&status).Error
	if err != nil {
		log.Println("failed to find status user with : ", err)
		return models.Status{}, errors.New("failed to find status with : " + err.Error())
	}

	if status.ID == 0 {
		log.Println("status user not found")
		return models.Status{}, errors.New("status user not found")
	}
	log.Println("status code : ", status.Code)
	return status, nil
}
