package user

import (
	"context"

	"github.com/phanphuctho7760/go-clean-architecture/app/entities"
	entitiesUtils "github.com/phanphuctho7760/go-clean-architecture/app/utils/entities"
)

type RepoItf interface {
	Create(ctx context.Context, u entities.User) (user entities.User, err entitiesUtils.ErrorCustom)
	GetById(ctx context.Context, id string) (user entities.User, err entitiesUtils.ErrorCustom)
	GetOneByField(ctx context.Context, fieldName string) (user entities.User, err entitiesUtils.ErrorCustom)
	GetMoreByField(ctx context.Context, fieldName string) (user []entities.User, err entitiesUtils.ErrorCustom)
	UpdateById(ctx context.Context, id string, u entities.User) (status bool, err entitiesUtils.ErrorCustom)
	DeleteById(ctx context.Context, id string) (status bool, err entitiesUtils.ErrorCustom)
}
