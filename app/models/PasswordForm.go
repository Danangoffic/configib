package models

import (
	"errors"
	"strings"
)

type PasswordForm struct {
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
	OBMParameter    map[string]string
}

func (p *PasswordForm) Clear() *PasswordForm {
	p.ConfirmPassword = ""
	p.NewPassword = ""
	p.OldPassword = ""
	return p
}

func (p *PasswordForm) Validate() error {
	if p.NewPassword == "" && p.OldPassword == "" {
		err := errors.New("Password cannot be empty")
		return err
	}
	if p.NewPassword != p.ConfirmPassword && strings.Compare(p.NewPassword, p.ConfirmPassword) != 0 {
		err := errors.New("Password does not match")
		return err
	}
	return nil
}
