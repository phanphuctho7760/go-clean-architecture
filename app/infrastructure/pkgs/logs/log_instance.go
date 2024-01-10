package logs

import (
	"log"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/colorconst"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

var LogGlobalInstance LogItf

func NewLogGlobalInstance(logLevel string) func() {
	var shutdownLogGlobalInstance func()
	LogGlobalInstance, shutdownLogGlobalInstance = newZapLogger("error")

	return func() {
		shutdownLogGlobalInstance()
		log.Printf("%sCall shutdown log global instance done at %s%s\n", colorconst.ANSIColorBlue, helpers.GetCallerLocationSkip(1), colorconst.ANSIColorWhite)
	}
}
