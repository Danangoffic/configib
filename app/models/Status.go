package models

type Status struct {
	ID          int64  `DB:"ID" json:"id"`
	Type        string `DB:"TYPE" json:"type"`
	Code        string `DB:"CODE" json:"code"`
	Name        string `DB:"NAME" json:"name"`
	Description string `DB:"DESCRIPTION" json:"description"`
}
