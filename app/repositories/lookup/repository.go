package lookup

import "github.com/Danangoffic/configib/app/models"

type Repository interface {
	FindByTypeAndCode(typeL, codeL string) (models.Lookup, error)
}
