package devicelockdown

import "github.com/Danangoffic/configib/app/models"

type Repository interface {
	FindByDeviceCode(deviceCode string) (models.DeviceLockdown, error)
}
