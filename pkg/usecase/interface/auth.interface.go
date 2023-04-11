package interfaces

import "context"

type AuthUseCase interface {
	SendOTP(ctx context.Context, phoneNumber string) error
	VarifyOTP(ctx context.Context, phoneNumber string, otp string) error
}
