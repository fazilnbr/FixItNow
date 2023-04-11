package usecase

import (
	"context"

	"github.com/fazilnbr/project-workey/pkg/config"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type authUseCase struct {
	adminRepo    interfaces.AdminRepository
	workerRepo   interfaces.WorkerRepository
	userRepo     interfaces.UserRepository
	mailConfig   config.MailConfig
	twilioConfig config.TwilioConfig
	config       config.Config
}

// SendOTP implements interfaces.AuthUseCase
func (c *authUseCase) SendOTP(ctx context.Context, phoneNumber string) error {
	return c.twilioConfig.SendOTP(c.config, phoneNumber)
}

// VarifyOTP implements interfaces.AuthUseCase
func (c *authUseCase) VarifyOTP(ctx context.Context, phoneNumber string, otp string) error {
	return c.twilioConfig.VerifyOTP(c.config, phoneNumber, otp)
}

func NewAuthService(
	adminRepo interfaces.AdminRepository,
	workerRepo interfaces.WorkerRepository,
	userRepo interfaces.UserRepository,
	mailConfig config.MailConfig,
	twilioConfig config.TwilioConfig,
	config config.Config,
) services.AuthUseCase {
	return &authUseCase{
		adminRepo:    adminRepo,
		workerRepo:   workerRepo,
		userRepo:     userRepo,
		mailConfig:   mailConfig,
		twilioConfig: twilioConfig,
		config:       config,
	}
}
