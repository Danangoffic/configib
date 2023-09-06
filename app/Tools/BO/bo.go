package bo

import (
	"github.com/Danangoffic/configib/app/repositories/lookup"
	"github.com/Danangoffic/configib/app/repositories/status"
	"github.com/Danangoffic/configib/app/repositories/systemparameter"
	"github.com/Danangoffic/configib/app/repositories/user"
)

type bo struct {
	statusRepo status.Repository
	userRepo user.Repository
	lookupRepo lookup.Repository
	systemparamRepo systemparameter.Repository
}

func NewBO(
	statusRepo status.Repository,
	userRepo user.Repository,
	lookupRepo lookup.Repository,
	systemparamRepo systemparameter.Repository,
) BO {
	return &bo{
		statusRepo: statusRepo,
		userRepo: userRepo,
		lookupRepo: lookupRepo,
		systemparamRepo: systemparamRepo,
	}
}