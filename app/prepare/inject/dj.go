package inject

import (
	"gorm.io/gorm"

	"github.com/phanphuctho7760/go-clean-architecture/app/adapters/controllers/user"
	handlersInfra "github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/handlers"
	userHandlersInfra "github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/handlers/user"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/logs"
	userRepoInfra "github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/repositories/user"
	userUc "github.com/phanphuctho7760/go-clean-architecture/app/usecases/user"
)

func InitDependService(
	db *gorm.DB,
) handlersInfra.AppHandlerItf {
	logger := logs.LogGlobalInstance

	// Init repository
	userRepo := userRepoInfra.NewRepo(
		db,
		logger,
	)

	// Init use case
	userUseCase := userUc.NewUserUseCase(
		userRepo,
	)

	// Init controller
	userController := user.NewUserController(
		userUseCase,
	)

	// Init handler
	userHandler := userHandlersInfra.NewHandler(userController)

	// Init use case
	appHandler := handlersInfra.NewAppHandler(userHandler)

	return appHandler
}
