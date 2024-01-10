package user

import (
	"context"

	"github.com/phanphuctho7760/go-clean-architecture/app/entities"
	entitiesUtils "github.com/phanphuctho7760/go-clean-architecture/app/utils/entities"
)

type UseCaseItf interface {
	Create(ctx context.Context, u entities.User) (user entities.User, err entitiesUtils.ErrorCustom)
	GetById(ctx context.Context, id string) (user entities.User, err entitiesUtils.ErrorCustom)
}
