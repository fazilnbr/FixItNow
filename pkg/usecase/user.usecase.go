package usecase

import (
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserService(
	userRepo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
