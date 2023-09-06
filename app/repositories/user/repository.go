package user

import "github.com/Danangoffic/configib/app/models"

type Repository interface {
	FindByID(id int64) (models.AppUser, error)
	FindByMobileNumberIncludeForceReset(loginName string) (models.AppUser, error)
}
