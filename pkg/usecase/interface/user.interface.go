package interfaces

import (
	"context"
)

type UserUseCase interface {
	RegisterAndVarifyWithNumber(ctx context.Context, phoneNumber string) (int, error)
	RegisterAndVarifyWithEmail(ctx context.Context, email string) (int, error)
}
