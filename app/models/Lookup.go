package models

type Lookup struct {
	ID          int64  `gorm:"column:ID" json:"id"`
	Type        string `gorm:"column:TYPE" json:"type"`
	Code        string `gorm:"column:CODE" json:"code"`
	Name        string `gorm:"column:NAME" json:"name"`
	Priority    int64  `gorm:"column:PRIORITY" json:"priority"`
	Description string `gorm:"column:DESCRIPTION" json:"description"`
	Shortname   string `gorm:"column:SHORTNAME" json:"shortname"`
	StatusID    int64  `gorm:"column:STATUS" json:"status_id"`
	Status      Status `gorm:"foreignKey:StatusID;references:ID" json:"status"`
	Filter      string `gorm:"column:FILTER" json:"filter"`
}
