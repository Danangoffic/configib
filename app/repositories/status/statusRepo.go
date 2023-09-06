package status

import (
	"errors"
	"log"

	"github.com/Danangoffic/configib/app/models"
	"gorm.io/gorm"
)

type statusRepo struct {
	DB *gorm.DB
}

func GetStatusRepository(db *gorm.DB) Repository {
	return &statusRepo{
		DB: db,
	}
}

func (r *statusRepo) FindByID(id int64) (models.Status, error) {
	var statusResut = models.Status{ID: id}

	err := r.DB.Unscoped().Find(&statusResut).Error
	if err != nil {
		log.Println("failed to find status data with : ", err)
		return models.Status{}, err
	}
	if statusResut.ID == 0 {
		log.Println("status data not found")
		return statusResut, errors.New("status not found")
	}
	return statusResut, nil
}
