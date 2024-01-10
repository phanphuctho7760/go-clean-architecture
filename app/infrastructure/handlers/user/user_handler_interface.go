package user

import (
	"github.com/gin-gonic/gin"
)

type HandlerItf interface {
	Create(c *gin.Context)
}
