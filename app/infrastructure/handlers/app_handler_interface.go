package handlers

import (
	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/handlers/user"
)

type AppHandlerItf interface {
	User() user.HandlerItf
}
