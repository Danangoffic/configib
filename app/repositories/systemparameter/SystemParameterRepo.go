package systemparameter

import (
	"errors"
	"log"

	"github.com/Danangoffic/configib/app/Tools"
	"github.com/Danangoffic/configib/app/models"
	"gorm.io/gorm"
)

type systemParamRepo struct {
	db *gorm.DB
}

func GetSystemParamRepository(db *gorm.DB) Repository {
	return &systemParamRepo{
		db: db,
	}
}

func (r *systemParamRepo) GetOnTheFly(vgroup, parameter string) (models.SystemParameter, error) {
	var systemParam models.SystemParameter

	err := r.db.Table("SYSTEM_PARAMETER").
		Where("1=1").
		Where(&models.SystemParameter{Vgroup: vgroup, Parameter: parameter}).
		Unscoped().
		First(&systemParam).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("not found : ", err)
			return models.SystemParameter{}, errors.New("not found")
		}
		if err == gorm.ErrInvalidData {
			log.Println("invalid data with ", err)
			return models.SystemParameter{}, errors.New("invalid")
		}
		log.Println("failed to get system param with : ", err)
		return models.SystemParameter{}, errors.New("failed")
	}
	appFact := Tools.NewAppFactory()
	systemParam.Svalue = appFact.DecryptData(systemParam.Svalue)
	return systemParam, nil
}

func (r *systemParamRepo) FindMaxFail() (models.SystemParameter, error) {
	var systemParam = models.SystemParameter{Parameter: "MAX_FAIL"}

	err := r.db.Unscoped().Table("SYSTEM_PARAMETER").Where(&systemParam).First(&systemParam).Error
	if err != nil {
		log.Println("failed to get system param with : ", err)
		return models.SystemParameter{}, errors.New("failed")
	}
	if systemParam.Svalue == "" {
		return systemParam, nil
	}
	appFact := Tools.NewAppFactory()
	systemParam.Svalue = appFact.DecryptData(systemParam.Svalue)
	return systemParam, nil
}
