package telemetries

import (
	"log"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/colorconst"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

func NewTelemetryGlobalConfig() (shutdown func()) {
	var shutdownTelemetryGlobalConfig func()

	name := "initOpenTelemetry"
	shutdownTelemetryGlobalConfig = initOpenTelemetry()

	return func() {
		shutdownTelemetryGlobalConfig()
		log.Printf("%sCall shutdown telemetry global instance %s done at %s%s\n", colorconst.ANSIColorBlue, name, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
	}
}

// NewTelemetry
// This function to create an instance for using - this is only can use after init
func NewTelemetry() TelemetryItf {
	return newOpenTelemetry()
}
