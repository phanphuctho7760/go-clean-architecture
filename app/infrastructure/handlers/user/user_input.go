package user

import (
	"context"

	"github.com/gin-gonic/gin"

	userInput "github.com/phanphuctho7760/go-clean-architecture/app/adapters/inputs/user"
	"github.com/phanphuctho7760/go-clean-architecture/app/entities"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/telemetries"
)

type createInput struct {
	c *gin.Context
}

func NewCreateInput(c *gin.Context) userInput.InputCreateItf {
	return &createInput{
		c: c,
	}
}

func (receiver createInput) Parse(ctx context.Context) (user entities.User) {
	name := "User.Input.Create"

	telemetry := telemetries.NewTelemetry()
	defer telemetry.EndSpan()

	telemetry.StartSpan(ctx, name)

	err := receiver.c.ShouldBindJSON(&user)
	if err != nil {
		return
	}
	return
}
