package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/colorconst"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/configs"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/error_tracking"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/logs"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/random_string"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/telemetries"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/uuids"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/validators"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/routers"
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/storage"
	"github.com/phanphuctho7760/go-clean-architecture/app/prepare/inject"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
	_ "github.com/phanphuctho7760/go-clean-architecture/docs"
)

//	@title			Go Clean Simple
//	@version		1.0
//	@description	This is a simple clean architect project
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@schemes	http https
//	@host		localhost:9999
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	limitWaitingTimeout := 3 * time.Second

	configs.NewConfigEnvGlobalInstance(helpers.GetProjectRootPath())
	uuids.NewUuidGlobalInstance()
	random_string.NewXidGlobalInstance()
	validators.NewValidatorGlobalInstance()

	shutdownErrorTrackingInstance := error_tracking.NewErrorTrackingGlobalInstance()
	defer shutdownErrorTrackingInstance()

	shutdownTelemetryGlobalConfig := telemetries.NewTelemetryGlobalConfig()
	defer shutdownTelemetryGlobalConfig()

	shutdownLogGlobalInstance := logs.NewLogGlobalInstance("debug")
	defer shutdownLogGlobalInstance()

	// Init db
	dbGorm := storage.NewDBGormPostgreSQL()
	err := dbGorm.Connect()
	if err != nil {
		// handler error here
	}

	// Add otel to gorm to get metric
	if err := dbGorm.DB.Use(otelgorm.NewPlugin()); err != nil {
		// handler error here
	}

	// Init repository, use case, controller and return handler
	appHandler := inject.InitDependService(
		dbGorm.DB,
	)

	// Init router
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(otelgin.Middleware(configs.ConfigEnvGlobalInstance.GetServiceName()))

	routers.InitRouter(
		router,
		appHandler,
	)

	srv := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		Addr:              ":" + configs.ConfigEnvGlobalInstance.GetPort(),
		Handler:           router,
	}

	go func() {
		log.Printf("Server serve at: %s\n\n\n", configs.ConfigEnvGlobalInstance.GetPort())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Fail to start server error %s at %s \n", err, helpers.GetCallerLocationSkip(1))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	sigChan := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't catch, so don't need to add it
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Printf("%sWaiting shutdown server ...%s", colorconst.ANSIColorCyan, colorconst.ANSIColorWhite)

	ctx, cancel := context.WithTimeout(context.Background(), limitWaitingTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("%sServer Shutdown error: %s at %s%s", colorconst.ANSIColorRed, err, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
	}
	// catching ctx.Done(). timeout of 3 seconds.
	select {
	case <-ctx.Done():
		log.Println("Timeout of 3 seconds at", helpers.GetCallerLocationSkip(1))
	}
	log.Printf("%sServer exited%s\n", colorconst.ANSIColorPurple, colorconst.ANSIColorWhite)
}
