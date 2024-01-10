package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/phanphuctho7760/go-clean-architecture/app/entities"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/error_tracking"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/telemetries"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/uuids"
	userRepos "github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/repositories/user"
	entitiesUtils "github.com/phanphuctho7760/go-clean-architecture/app/utils/entities"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

type UseCase struct {
	repo userRepos.RepoItf
}

func NewUserUseCase(
	r userRepos.RepoItf,
) UseCaseItf {
	return &UseCase{
		repo: r,
	}
}

func (receiver UseCase) Create(ctx context.Context, u entities.User) (user entities.User, err entitiesUtils.ErrorCustom) {
	name := "User.UseCase.Create"

	telemetry := telemetries.NewTelemetry()
	newCtx := telemetry.StartSpan(ctx, name)
	defer telemetry.EndSpan()

	newCtx = helpers.MakeTrackIDIntoNewContext(ctx, newCtx)

	fmt.Printf("SpanID UC Before: %s\n", telemetry.GetSpanID())

	u.Id = uuids.UuidInstance.GenerateUUIDString()

	user, err = receiver.repo.Create(newCtx, u)

	fmt.Printf("SpanID UC After: %s\n", telemetry.GetSpanID())

	if err.Inner != nil {

		telemetry.RecordError(newCtx, err.Inner)
		telemetry.SetStatusError(newCtx, err.Inner.Error())
		uStr, _ := json.Marshal(u)
		telemetry.SetAttributeString(newCtx, "uFromUCSpan", string(uStr))

		error_tracking.ErrorTrackingGlobalInstance.SetExtra(newCtx, "uFromUC", u)
		error_tracking.ErrorTrackingGlobalInstance.CaptureError(newCtx, err)

		return
	}

	telemetry.SetStatusSuccess(ctx, "ok")

	return
}

func (receiver UseCase) GetById(ctx context.Context, id string) (user entities.User, err entitiesUtils.ErrorCustom) {
	name := "User.UseCase.GetById"

	telemetry := telemetries.NewTelemetry()
	newCtx := telemetry.StartSpan(ctx, name)
	defer telemetry.EndSpan()

	newCtx = helpers.MakeTrackIDIntoNewContext(ctx, newCtx)

	user, err = receiver.repo.GetById(ctx, id)

	if err.Inner != nil {
		telemetry.RecordError(newCtx, err.Inner)
		telemetry.SetStatusError(newCtx, err.Inner.Error())
		telemetry.SetAttributeString(newCtx, "idFromUCSpan", id)

		error_tracking.ErrorTrackingGlobalInstance.SetExtra(newCtx, "uFromUC", id)
		error_tracking.ErrorTrackingGlobalInstance.CaptureError(newCtx, err)

		return
	}

	return
}
