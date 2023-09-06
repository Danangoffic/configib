package systemparameter

import "github.com/Danangoffic/configib/app/models"

type Repository interface {
	GetOnTheFly(vgroup, parameter string) (models.SystemParameter, error)
	FindMaxFail() (models.SystemParameter, error)
}
