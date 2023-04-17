package interfaces

import (
	"context"

	"github.com/fazilnbr/project-workey/pkg/domain"
)

type UserUseCase interface {
	RegisterAndVarifyWithNumber(ctx context.Context, phoneNumber string) (int, error)
	RegisterAndVarifyWithEmail(ctx context.Context, email string) (int, error)
	AddProfile(ctx context.Context, userData domain.UserData) error
	UpdateMail(ctx context.Context, email string, userId int) error
	GetProfile(ctx context.Context, userId int) (domain.Profile, error)
}
