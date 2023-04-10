package handler

import (
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type UserHandler struct {
	userService services.UserUseCase
}

func NewUserHandler(userService services.UserUseCase) UserHandler {
	return UserHandler{
		userService: userService,
	}
}
