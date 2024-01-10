package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/random_string"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/constants"
)

// TrackMiddleWare
// Add track id to request context
// This should only use in middleware
func TrackMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constants.TrackIDKey, random_string.XidGlobalInstance.Generate20Character())
		c.Next()
	}
}
