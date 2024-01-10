package error_tracking

import (
	"log"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/colorconst"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

var ErrorTrackingGlobalInstance ErrorTrackingItf

func NewErrorTrackingGlobalInstance() (shutdown func()) {
	var shutdownErrorTrackingGlobalInstance func()
	ErrorTrackingGlobalInstance, shutdownErrorTrackingGlobalInstance = newSentryTrackingWithOpenTelemetry()

	return func() {
		shutdownErrorTrackingGlobalInstance()
		log.Printf("%sCall shutdown error tracking global instance done at %s%s\n", colorconst.ANSIColorBlue, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
	}
}
