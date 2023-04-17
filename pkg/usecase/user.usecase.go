package usecase

import (
	"context"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

// UpdateMail implements interfaces.UserUseCase
func (c *userUseCase) UpdateMail(ctx context.Context, email string, userId int) error {
	err := c.userRepo.UpdateMail(ctx, email, userId)
	return err
}

// AddProfile implements interfaces.UserUseCase
func (c *userUseCase) AddProfile(ctx context.Context, userData domain.UserData) error {
	err := c.userRepo.AddProfile(ctx, userData)
	return err
}

// RegisterAndVarifyWithEmail implements interfaces.UserUseCase
func (c *userUseCase) RegisterAndVarifyWithEmail(ctx context.Context, email string) (int, error) {
	user, err := c.userRepo.FindUserWithEmail(ctx, email)
	if err == nil || err.Error() != "there is no user" {
		return user.IdUser, err
	}
	id, err := c.userRepo.CreateUser(ctx, domain.User{
		Email: email,
		Phone: utils.Randomphone(5),
	})
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
		Phone: phoneNumber,
		Email: utils.Randommail(5),
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
