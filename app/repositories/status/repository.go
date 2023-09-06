package status

import "github.com/Danangoffic/configib/app/models"

type Repository interface {
	FindByID(id int64) (models.Status, error)
}
