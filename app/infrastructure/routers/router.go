package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/handlers"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/middlewares"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/migration"
)

func InitRouter(
	r *gin.Engine,
	appHandler handlers.AppHandlerItf,
) *gin.Engine {
	apiV1 := r.Group("api/v1", middlewares.TrackMiddleWare())
	{
		usersV1 := apiV1.Group("user")
		{
			usersV1.POST("", appHandler.User().Create)
		}

		apiV1.POST("migrate", func(c *gin.Context) {
			// TODO make private call by ssh,...
			key := c.PostForm("key")
			if key == "yourkeyhere" {
				migration.RunMigrateGorm()
				c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
				return
			}

			c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "failed"})
		})

	}
	return r
}
