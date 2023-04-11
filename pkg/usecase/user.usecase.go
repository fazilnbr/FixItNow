package usecase

import (
	"context"
	"fmt"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

// RegisterAndVarifyWithEmail implements interfaces.UserUseCase
func (c *userUseCase) RegisterAndVarifyWithEmail(ctx context.Context, email string) (int, error) {
	user, err := c.userRepo.FindUserWithEmail(ctx, email)
	if err == nil || err.Error() != "there is no user" {
		return user.IdUser, err
	}
	id, err := c.userRepo.CreateUser(ctx, domain.User{
		Email:    email,
		Phone:    utils.Randomphone(5),
		UserName: utils.RandomString(5),
	})
	fmt.Printf("\n\nerr : %v\n\n", err)
	if err != nil {
		return 0, err
	}
	return id, err
}

// RegisterAndVarify implements interfaces.UserUseCase
func (c *userUseCase) RegisterAndVarifyWithNumber(ctx context.Context, phoneNumber string) (int, error) {
	user, err := c.userRepo.FindUserWithNumber(ctx, phoneNumber)
	if err == nil || err.Error() != "there is no user" {
		return user.IdUser, err
	}
	id, err := c.userRepo.CreateUser(ctx, domain.User{
		Phone:    phoneNumber,
		Email:    utils.Randommail(5),
		UserName: utils.RandomString(5),
	})
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
