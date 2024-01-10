package handlers

import (
	userHandlers "github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/handlers/user"
)

type appHandler struct {
	user userHandlers.HandlerItf
}

func NewAppHandler(
	user userHandlers.HandlerItf,
) AppHandlerItf {
	app := appHandler{
		user: user,
	}
	return &app
}

func (receiver appHandler) User() userHandlers.HandlerItf {
	return receiver.user
}
