package handlers

import (
	"net/http"
	"regexp"

	// "github.com/Danangoffic/configib/app/usecase"
	"github.com/Danangoffic/configib/app/models"
	"github.com/gin-gonic/gin"
)

// type LoginHandler struct {
// 	LoginUsecase usecase.Usecase
// }

// func NewLoginHandler(usecase usecase.Usecase) *handler {
// 	return &LoginHandler{
// 		LoginUsecase: usecase
// 	}
// }

func (h *HandlerModel) DoLoginScrumV2(c *gin.Context) {
	var req models.LoginRequestScrum
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.LoginUsecase.DoLoginScrumV2(c)
	return
}

func isNum(s string) bool {
	if s == "" {
		return false
	}

	matched, err := regexp.MatchString("[-+]?\\d*\\.?\\d+", s)
	if err != nil {
		return false
	}

	return matched
}
