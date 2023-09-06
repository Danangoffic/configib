package account

import (
	"errors"
	"log"

	"github.com/Danangoffic/configib/app/models"
	"gorm.io/gorm"
)

type accountRepo struct {
	DB *gorm.DB
}

func GetAccountRepository(db *gorm.DB) Repository {
	return &accountRepo{
		DB: db,
	}
}
func (r *accountRepo) Get(id int64) (models.Account, error) {
	var account = models.Account{ID: id}

	err := r.DB.Unscoped().Find(&account).Error
	if err != nil {
		log.Println("failed to find account with : ", err)
	}
	return account, nil
}

func (r *accountRepo) FindByCustomerID(id int64) ([]models.Account, error) {
	var accounts []models.Account
	err := r.DB.Unscoped().Preload("Customer").
		Where("IS_HIDDEN = ?", "NO").
		Where("CUSTOMER_ID = ?", id).
		Where("STATUS = (SELECT * FROM STATUS WHERE TYPE = 'account' AND CODE = 'active')").
		Find(&accounts).Error
	if err != nil {
		log.Println("failed to find accounts with : ", err)
		return []models.Account{}, err
	}
	if len(accounts) == 0 {
		log.Println("account not found")
		return []models.Account{}, errors.New("account not found")
	}
	return accounts, nil

}

func (r *accountRepo) GetByAccountNumber(cifCode interface{}, accountNumber string) (models.Account, error) {
	var account = models.Account{AccountNumber: accountNumber}
	sql := r.DB.Unscoped().
		Where("STATUS IN ",
			r.DB.Table("STATUS").
				Select("ID").
				Where("CODE", "active").
				Or("CODE", "dormant").
				Or("CODE", "col2")).
		Or(r.DB.Where("ACCOUNT_TYPE", "VirtualCreditCardAccount").
			Where("STATUS", r.DB.Table("STATUS").
				Select("ID").
				Where("CODE", "notActivated")))
	if cifCode != nil || cifCode.(string) != "" {
		sql.Where("CUSTOMER_ID", r.DB.
			Table("CUSTOMER").
			Select("ID").
			Where("CIF_CODE = ?", cifCode.(string)))
	}
	err := sql.Find(&account).Error
	if err != nil {
		log.Println("failed to find account with : ", err)
	}
	return account, nil
}
