package user

import (
	"context"

	entitiesUtils "github.com/phanphuctho7760/go-clean-architecture/app/utils/entities"
)

type (
	CreateItf interface {
		Present(ctx context.Context, statusCode entitiesUtils.StatusCode, data any) any
	}
)
