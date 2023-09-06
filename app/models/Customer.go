package models

import (
	"time"
)

type Customer struct {
	ID                int64     `gorm:"primaryKey;autoIncrement;column:ID"`
	CifCode           string    `gorm:"column:CIF_CODE"`
	Level             int64     `gorm:"column:LEVEL"`
	Type              int64     `gorm:"column:TYPE"`
	Title             string    `gorm:"column:TITLE"`
	Name              string    `gorm:"column:NAME"`
	Email             string    `gorm:"column:EMAIL"`
	StatusID          int64     `gorm:"column:STATUS"`
	Status            Status    `gorm:"foreignKey:StatusID;references:ID"`
	CreatedDate       time.Time `gorm:"column:CREATED_DATE"`
	CreatedBy         string    `gorm:"column:CREATED_BY"`
	LastUpdatedDate   time.Time `gorm:"column:LAST_UPDATED_DATE"`
	LastUpdatedBy     string    `gorm:"column:LAST_UPDATED_BY"`
	CompanyCode       string    `gorm:"column:COMPANY_CODE"`
	EmployersCode     string    `gorm:"column:EMPLOYERS_CODE"`
	Gender            string    `gorm:"column:GENDER"`
	BirthDate         time.Time `gorm:"column:BIRTH_DATE"`
	MothersMaiden     string    `gorm:"column:MOTHERS_MAIDEN"`
	IdCardNumber      string    `gorm:"column:ID_CARD_NUMBER"`
	CifCodeNKYC       string    `gorm:"column:CIF_CODE_NKYC"`
	CustomerSegmentID int64     `gorm:"column:CUSTOMER_SEGMENT"`
	CustomerSegment   Lookup    `gorm:"foreignKey:CustomerSegmentID"`
	Accounts          []Account `gorm:"foreignKey:CUSTOMER_ID;references:ID"`
}
