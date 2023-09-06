package handlers

import (
	"github.com/Danangoffic/configib/app/usecase/login"
	"github.com/gin-gonic/gin"
)

type handler interface {
	DoLoginScrumV2(c *gin.Context)
}

type HandlerModel struct {
	LoginUsecase login.Usecase
}

func NewHandler(h HandlerModel) handler {
	return &h
}
