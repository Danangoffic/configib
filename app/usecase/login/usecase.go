package login

import "github.com/gin-gonic/gin"

type Usecase interface {
	DoLoginScrumV2(c *gin.Context)
}
