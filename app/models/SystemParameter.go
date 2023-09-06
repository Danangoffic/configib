package models

import "time"

type SystemParameter struct {
	ID              int64     `gorm:"column:ID"`
	Vgroup          string    `gorm:"column:VGROUP"`
	Parameter       string    `gorm:"column:PARAMETER"`
	Svalue          string    `gorm:"column:SVALUE"`
	Description     string    `gorm:"column:DESCRIPTION"`
	StatusID        string    `gorm:"column:STATUS"`
	CreatedDate     time.Time `gorm:"column:CREATED_DATE"`
	CreatedBy       string    `gorm:"column:CREATED_BY"`
	LastUpdatedDate time.Time `gorm:"column:LAST_UPDATED_DATE"`
	LastUpdatedBy   string    `gorm:"column:LAST_UPDATED_BY"`
	UserLevel       int64     `gorm:"column:USER_LEVEL"`
}
