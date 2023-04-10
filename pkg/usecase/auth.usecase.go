package usecase

import (
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
