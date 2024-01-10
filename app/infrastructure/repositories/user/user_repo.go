package user

import (
	"context"
	"encoding/json"
	"runtime/debug"

	"gorm.io/gorm"

	"github.com/phanphuctho7760/go-clean-architecture/app/entities"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/error_tracking"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/logs"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/telemetries"
	entitiesUtils "github.com/phanphuctho7760/go-clean-architecture/app/utils/entities"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

type Repo struct {
	db     *gorm.DB
	logger logs.LogItf
}

func NewRepo(
	db *gorm.DB,
	logger logs.LogItf,
) RepoItf {
	return &Repo{
		db:     db,
		logger: logger,
	}
}

func (receiver Repo) Create(ctx context.Context, u entities.User) (user entities.User, err entitiesUtils.ErrorCustom) {
	name := "User.Repo.Create"
	trackId := helpers.GetTrackIDFromContext(ctx)

	telemetry := telemetries.NewTelemetry()
	defer telemetry.EndSpan()
	telemetry.StartSpan(ctx, name)

	receiver.logger.Debugf("SpanID Repo Before: %s\n", telemetry.GetSpanID())

	tx := receiver.db.WithContext(ctx).Create(&u)

	receiver.logger.Debugf("SpanID Repo After: %s\n", telemetry.GetSpanID())

	if tx.Error != nil {
		// Remove ID when error
		u.Id = ""

		telemetry.RecordError(ctx, tx.Error)
		telemetry.SetStatusError(ctx, tx.Error.Error())
		uStr, _ := json.Marshal(u)
		telemetry.SetAttributeString(ctx, "uFromRepoSpan", string(uStr))

		error_tracking.ErrorTrackingGlobalInstance.CaptureError(ctx, err)

		err = entitiesUtils.ErrorCustom{
			Inner:      tx.Error,
			Message:    "Error create new user",
			StackTrace: string(debug.Stack()),
			Misc: map[string]interface{}{
				"user": u,
			},
		}

		receiver.logger.Errorf("Error create user. track_id: %s", trackId, tx.Error)

		return u, err
	}

	telemetry.SetStatusSuccess(ctx, "ok")

	return u, entitiesUtils.ErrorCustom{}
}

func (receiver Repo) GetById(ctx context.Context, id string) (user entities.User, err entitiesUtils.ErrorCustom) {
	// TODO implement me
	panic("implement me")
}

func (receiver Repo) GetOneByField(ctx context.Context, fieldName string) (user entities.User, err entitiesUtils.ErrorCustom) {
	// TODO implement me
	panic("implement me")
}

func (receiver Repo) GetMoreByField(ctx context.Context, fieldName string) (user []entities.User, err entitiesUtils.ErrorCustom) {
	// TODO implement me
	panic("implement me")
}

func (receiver Repo) UpdateById(ctx context.Context, id string, u entities.User) (status bool, err entitiesUtils.ErrorCustom) {
	// TODO implement me
	panic("implement me")
}

func (receiver Repo) DeleteById(ctx context.Context, id string) (status bool, err entitiesUtils.ErrorCustom) {
	// TODO implement me
	panic("implement me")
}
