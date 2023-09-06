package models

type LoginHistory struct {
	ID     int64 `DB:"ID" json:"id"`
	UserID int64 `json:"userId"`
	// User           User        `json:"user"`
	LoginTime      interface{} `DB:"LOGIN_TIME" json:"loginTime"`
	LogoutTime     interface{} `json:"logoutTime"`
	LastAccess     interface{} `DB:"LAST_ACCESS" json:"lastAccess"`
	Status         Status      `gorm:"references:ID" json:"status"`
	SessionID      string      `DB:"SESSION_ID" json:"sessionId"`
	Description    string      `DB:"DESCRIPTION" json:"description"`
	DeviceLockDown DeviceLockdown
}
