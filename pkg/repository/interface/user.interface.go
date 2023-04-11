package interfaces

import (
	"context"

	"github.com/fazilnbr/project-workey/pkg/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (int, error)
	FindUserWithNumber(ctx context.Context, phoneNumber string) (domain.User, error)
	FindUserWithEmail(ctx context.Context, email string) (domain.User, error)
}
