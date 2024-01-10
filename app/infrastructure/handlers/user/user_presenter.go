package user

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phanphuctho7760/go-clean-architecture/app/entities"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/telemetries"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/constants"
	entitiesUtils "github.com/phanphuctho7760/go-clean-architecture/app/utils/entities"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

type CreatePresenter struct {
	c *gin.Context
}

func NewCreatePresenter(c *gin.Context) CreatePresenter {
	return CreatePresenter{
		c: c,
	}
}

func (receiver CreatePresenter) Present(ctx context.Context, statusCode entitiesUtils.StatusCode, data any) any {
	name := "User.Presenter.Create"

	telemetry := telemetries.NewTelemetry()
	telemetry.StartSpan(ctx, name)
	defer telemetry.EndSpan()

	trackId := helpers.GetTrackIDFromContext(ctx)
	receiver.c.Header(constants.TrackIDKey, trackId)

	if statusCode.GetCode() == entitiesUtils.SuccessStatusCode.GetCode() {
		receiver.c.JSON(http.StatusOK, data)
		return nil
	} else {
		receiver.c.JSON(statusCode.GetHTTPCode(), data)
		return nil
	}
}

type GetPresenter struct {
	c *gin.Context
}

func NewGetPresenter(c *gin.Context) GetPresenter {
	return GetPresenter{
		c: c,
	}
}

func (receiver GetPresenter) Present(users []entities.User) any {
	receiver.c.JSON(http.StatusOK, users)
	return nil
}
