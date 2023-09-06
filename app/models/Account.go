package models

import (
	"time"
)

type Account struct {
	ID                int64      `gorm:"column:ID;primaryKey;autoIncrement"`
	Name              string     `gorm:"column:NAME"`
	AccountType       string     `gorm:"column:ACCOUNT_TYPE"`
	ProductType       string     `gorm:"column:PRODUCT_TYPE"`
	AccountNumber     string     `gorm:"column:ACCOUNT_NUMBER"`
	Bank              string     `gorm:"column:BANK"`
	Notes             string     `gorm:"column:NOTES"`
	Currency          string     `gorm:"column:CURRENCY"`
	StatusID          int64      `gorm:"column:STATUS"`
	Status            Status     `gorm:"foreignKey:StatusID;references:ID"`
	CreatedDate       time.Time  `gorm:"column:CREATED_DATE"`
	CreatedBy         string     `gorm:"column:CREATED_BY"`
	LastUpdatedDate   time.Time  `gorm:"column:LAST_UPDATED_DATE"`
	LastUpdatedBy     string     `gorm:"column:LAST_UPDATED_BY"`
	AllowFlag         bool       `gorm:"column:ALLOW_FLAG"`
	ReedemFlag        bool       `gorm:"column:REEDEM_FLAG"`
	IsDefault         bool       `gorm:"column:IS_DEFAULT"`
	LinkingCard       string     `gorm:"column:LINKING_CARD"`
	IsAlreadyClose    bool       `gorm:"column:IS_ALREADY_CLOSE"`
	InitialDeposit    float64    `gorm:"column:INITIAL_DEPOSIT"`
	InitialDepositStr string     `gorm:"column:INITIAL_DEPOSIT_STR"`
	CustomerID        int64      `gorm:"column:CUSTOMER_ID"`
	Customer          Customer   `gorm:"foreignKey:CustomerID;references:ID"`
	BankBranchID      int64      `gorm:"column:BANK_BRANCH"`
	BankBranch        BankBranch `gorm:"foreignKey:BankBranchID;references:ID"`
	AccountTypeCode   string     `gorm:"column:ACCOUNT_TYPE_CODE"`
	MerchantCode      string     `gorm:"column:MERCHANT_CODE"`
	Label             string     `gorm:"column:LABEL"`
	AccountMode       string     `gorm:"column:ACCOUNT_MODE"`
	BranchCode        string     `gorm:"column:BRANCH_CODE"`
	ExpiryDate        string     `gorm:"column:EXPIRY_DATE"`
	CardType          string     `gorm:"column:CARD_TYPE"`
	IsHidden          bool       `gorm:"column:IS_HIDDEN"`
	RealAccount       string     `gorm:"column:REAL_ACCOUNT"`
	SharingAccount    string     `gorm:"column:SHARING_ACCOUNT"`
}
