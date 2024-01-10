package user

import (
	"context"
	"encoding/json"
	"fmt"

	userInput "github.com/phanphuctho7760/go-clean-architecture/app/adapters/inputs/user"
	userPresenters "github.com/phanphuctho7760/go-clean-architecture/app/adapters/presenters/user"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/error_tracking"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/telemetries"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/validators"
	"github.com/phanphuctho7760/go-clean-architecture/app/usecases/user"
	entitiesUtils "github.com/phanphuctho7760/go-clean-architecture/app/utils/entities"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

type userController struct {
	uc user.UseCaseItf
}

func NewUserController(
	uc user.UseCaseItf,
) ControllerItf {
	return &userController{
		uc: uc,
	}
}

// Create godoc
//
//	@Summary		Create user
//	@Description	create one user
//	@Accept			json
//	@Produce		json
//	@Param			Body	body		entities.User									true	"User Info"
//	@Success		200		{object}	entities.Response{data=entities.User}			"success"
//	@Failure		400		{object}	entities.Response{data=entities.errorFormat}	"fail"
//	@Router			/user [post]
func (receiver userController) Create(ctx context.Context, i userInput.InputCreateItf, p userPresenters.CreateItf) any {
	name := "User.Controller.Create"

	telemetry := telemetries.NewTelemetry()
	defer telemetry.EndSpan()

	newCtx := telemetry.StartSpan(ctx, name)
	newCtx = helpers.MakeTrackIDIntoNewContext(ctx, newCtx)

	userParam := i.Parse(newCtx)

	// parse param not validate data input, required validate data after parse
	err := validators.ValidatorGlobalInstance.ValidateStructOneErr(userParam)
	if err != nil {

		telemetry.RecordError(newCtx, err)
		telemetry.SetStatusError(newCtx, err.Error())
		userParamStr, _ := json.Marshal(userParam)
		telemetry.SetAttributeString(newCtx, "userParamFromCtlSpan", string(userParamStr))

		error_tracking.ErrorTrackingGlobalInstance.CaptureError(newCtx, err)

		statusBadRequest := entitiesUtils.BadRequestStatusCode
		data := entitiesUtils.MakeResponseErrorFormat(
			statusBadRequest.GetMessage(),
			statusBadRequest.GetCode(),
			err.Error(),
		)

		p.Present(newCtx, statusBadRequest, data)
		return nil
	}

	fmt.Printf("SpanID Ctl Before: %s\n", telemetry.GetSpanID())

	userNew, errCus := receiver.uc.Create(newCtx, userParam)

	fmt.Printf("SpanID Ctl After: %s\n", telemetry.GetSpanID())

	if errCus.Inner != nil {

		telemetry.RecordError(newCtx, errCus.Inner)
		telemetry.SetStatusError(newCtx, errCus.Inner.Error())
		userParamStr, _ := json.Marshal(userParam)
		telemetry.SetAttributeString(newCtx, "userParamFromCtlSpan", string(userParamStr))

		error_tracking.ErrorTrackingGlobalInstance.SetExtra(newCtx, "userParamFromCtlSpan", userParam)
		error_tracking.ErrorTrackingGlobalInstance.CaptureError(newCtx, errCus.Inner)

		internalServerErrorStatusCode := entitiesUtils.InternalServerErrorStatusCode
		data := entitiesUtils.MakeResponseErrorFormat(
			internalServerErrorStatusCode.GetMessage(),
			internalServerErrorStatusCode.GetCode(),
			"Have an error when create user",
		)

		p.Present(newCtx, internalServerErrorStatusCode, data)
		return nil
	}

	statusCodeSuccess := entitiesUtils.SuccessStatusCode
	data := entitiesUtils.MakeResponseSuccessFormat(
		statusCodeSuccess.GetMessage(),
		userNew,
	)

	p.Present(newCtx, statusCodeSuccess, data)
	return nil
}
