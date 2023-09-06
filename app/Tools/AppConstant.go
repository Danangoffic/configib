package Tools

import (
	"time"

	"github.com/Danangoffic/configib/app/models"
)

var CoreBusinessDate time.Time = time.Now()
var TxIdFile string = "txid.counter"

type appConstant struct {
}

type AppConstant interface {
	GetSystemParamByObj(obj models.SystemParameter) string
}

func NewAppConstant() AppConstant {
	return &appConstant{}
}

func (r *appConstant) GetSystemParamByObj(obj models.SystemParameter) string {
	return obj.Svalue
}
