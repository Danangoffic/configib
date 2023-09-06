package models

import "time"

type DeviceLockdown struct {
	ID              int64     `gorm:"ID" json:"id"`
	UserID          int64     `gorm:"column:USER_ID"`
	AppUser         AppUser   `gorm:"foreignKey:UserID;references:APP_USER(ID)"`
	DeviceCode      string    `gorm:"column:DEVICE_CODE" json:"deviceCode"`
	DeviceParameter string    `gorm:"column:DEVICE_PARAMETER" json:"deviceParameter"`
	DeviceInfo      string    `gorm:"column:DEVICE_INFO" json:"deviceInfo"`
	Status          Status    `gorm:"foreignKey:STATUS;references:ID" json:"status"`
	CreatedDate     time.Time `gorm:"column:CREATED_DATE" json:"createdDate"`
	CreatedBy       string    `gorm:"column:CREATED_BY" json:"createdBy"`
	LastUpdatedDate time.Time `gorm:"column:LAST_UPDATED_DATE" json:"lastUpdatedDate"`
	LastUpdatedBy   string    `gorm:"column:LAST_UPDATED_BY" json:"lastUpdatedBy"`
	IsLogin         bool
}
