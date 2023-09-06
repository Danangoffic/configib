package customer

import (
	"errors"
	"log"

	"github.com/Danangoffic/configib/app/models"
	"github.com/Danangoffic/configib/app/repositories/account"
	"github.com/Danangoffic/configib/app/repositories/status"
	"gorm.io/gorm"
)

type customerRepo struct {
	DB          *gorm.DB
	statusRepo  status.Repository
	accountRepo account.Repository
}

func GetCustomerRepository(db *gorm.DB, statusRepo status.Repository, accountRepo account.Repository) Repository {
	return &customerRepo{
		db,
		statusRepo,
		accountRepo,
	}
}

func (r *customerRepo) FindByID(id int64) (models.Customer, error) {
	var customer = models.Customer{ID: id}

	err := r.DB.Unscoped().Preload("Status").Find(&customer).Error
	if err != nil {
		log.Println("failed to find customer with : ", err)
		return models.Customer{}, err
	}

	if customer.ID == 0 {
		log.Println("customer data not found")
		return models.Customer{}, errors.New("customer data not found")
	}

	if customer.Status.ID == 0 {
		customer.Status, _ = r.statusRepo.FindByID(customer.StatusID)
	}

	if len(customer.Accounts) == 0 {
		customer.Accounts, _ = r.accountRepo.FindByCustomerID(customer.ID)
	}

	return customer, nil
}
