package account

import "github.com/Danangoffic/configib/app/models"

type Repository interface {
	Get(id int64) (models.Account, error)
	FindByCustomerID(id int64) ([]models.Account, error)
}
