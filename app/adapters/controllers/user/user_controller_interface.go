package user

import (
	"context"

	"github.com/phanphuctho7760/go-clean-architecture/app/adapters/inputs/user"
	userPresenters "github.com/phanphuctho7760/go-clean-architecture/app/adapters/presenters/user"
)

type ControllerItf interface {
	Create(ctx context.Context, i user.InputCreateItf, p userPresenters.CreateItf) any
}
