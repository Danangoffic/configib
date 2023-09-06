package models

import (
	"time"
)

type BankBranch struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`
	Code            string    `gorm:"column:CODE"`
	Mnemonic        string    `gorm:"column:MNEMONIC"`
	Name            string    `gorm:"column:NAME"`
	Address         string    `gorm:"column:ADDRESS"`
	CreatedDate     time.Time `gorm:"column:CREATED_DATE"`
	CreatedBy       string    `gorm:"column:CREATED_BY"`
	LastUpdatedDate time.Time `gorm:"column:LAST_UPDATED_DATE"`
	LastUpdatedBy   string    `gorm:"column:LAST_UPDATED_BY"`
	FilterCode      string    `gorm:"column:FILTER_CODE"`
}
