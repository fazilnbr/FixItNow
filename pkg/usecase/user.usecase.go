package usecase

import (
	"context"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

// RegisterAndVarify implements interfaces.UserUseCase
func (c *userUseCase) RegisterAndVarify(ctx context.Context, phoneNumber string) (int, error) {
	user, err := c.userRepo.FindUserWithNumber(ctx, phoneNumber)
	if err == nil || err.Error() != "there no user" {
		return user.IdUser, err
	}
	id, err := c.userRepo.CreateUser(ctx, domain.User{Phone: phoneNumber})
	if err != nil {
		return 0, err
	}
	return id, err
}

func NewUserService(
	userRepo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
