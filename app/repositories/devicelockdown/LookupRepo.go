package devicelockdown

import (
	"errors"
	"log"

	"github.com/Danangoffic/configib/app/models"
	"gorm.io/gorm"
)

type devicelockdownRepo struct {
	db *gorm.DB
}

func GetDeviceLockdownRepository(db *gorm.DB) Repository {
	return &devicelockdownRepo{
		db: db,
	}
}

func (r *devicelockdownRepo) FindByDeviceCode(deviceCode string) (models.DeviceLockdown, error) {
	var dld = models.DeviceLockdown{DeviceCode: deviceCode}
	log.Println("deviceCode : ", deviceCode)

	err := r.db.
		Table("DEVICE_LOCKDOWN").Preload("AppUser").
		Unscoped().
		Where("STATUS", r.db.Table("STATUS").Select("ID").Where("CODE", "active")).
		Find(&dld).Error
	if err != nil {
		log.Println("failed to find lookup data with : ", err)
		return models.DeviceLockdown{}, err
	}
	if dld.ID == 0 {
		log.Println("lookup data not found ")
		return models.DeviceLockdown{}, errors.New("not found")
	}
	// var user = models.AppUser{ID: dld.UserID}
	return dld, nil
}
