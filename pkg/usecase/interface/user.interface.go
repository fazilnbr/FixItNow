package interfaces

import (
	"context"
)

type UserUseCase interface {
	RegisterAndVarify(ctx context.Context, phoneNumber string) (int, error)
}
