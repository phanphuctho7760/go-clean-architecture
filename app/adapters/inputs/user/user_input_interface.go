package user

import (
	"context"

	"github.com/phanphuctho7760/go-clean-architecture/app/entities"
)

type (
	InputCreateItf interface {
		Parse(ctx context.Context) (user entities.User)
	}
)
