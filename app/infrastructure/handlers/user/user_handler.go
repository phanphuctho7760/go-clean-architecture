package user

import (
	"github.com/gin-gonic/gin"

	userControllers "github.com/phanphuctho7760/go-clean-architecture/app/adapters/controllers/user"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/telemetries"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

type handler struct {
	uCtl userControllers.ControllerItf
}

func NewHandler(
	uCtl userControllers.ControllerItf,
) HandlerItf {
	h := handler{
		uCtl: uCtl,
	}
	return &h
}

func (receiver handler) Create(c *gin.Context) {
	name := "User.Handler.Create"

	telemetry := telemetries.NewTelemetry()
	defer telemetry.EndSpan()

	newCtx := telemetry.StartSpan(c.Request.Context(), name)
	newCtx = helpers.MakeTrackIDIntoNewContext(c, newCtx)

	input := NewCreateInput(c)
	presenter := NewCreatePresenter(c)

	receiver.uCtl.Create(newCtx, input, presenter)
	return
}
