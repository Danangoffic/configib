package lookup

import (
	"errors"
	"log"

	"github.com/Danangoffic/configib/app/models"
	"gorm.io/gorm"
)

type lookupRepo struct {
	db *gorm.DB
}

func GetLookupRepository(db *gorm.DB) Repository {
	return &lookupRepo{
		db: db,
	}
}

func (r *lookupRepo) FindByTypeAndCode(typeL, codeL string) (models.Lookup, error) {
	var lk models.Lookup
	log.Println("type : ", typeL)
	log.Println("code : ", codeL)

	err := r.db.Table("LOOKUP").Unscoped().Where(models.Lookup{Type: typeL, Code: codeL}).Find(&lk).Error
	if err != nil {
		log.Println("failed to find lookup data with : ", err)
		return models.Lookup{}, err
	}
	if lk.ID == 0 {
		log.Println("lookup data not found ")
		return models.Lookup{}, errors.New("not found")
	}
	return lk, nil
}
