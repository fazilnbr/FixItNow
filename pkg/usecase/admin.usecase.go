package usecase

import (
	"github.com/fazilnbr/project-workey/pkg/config"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type adminUseCase struct {
	adminRepo  interfaces.AdminRepository
	mailConfig config.MailConfig
}

func NewAdminService(
	adminRepo interfaces.AdminRepository,
	mailConfig config.MailConfig) services.AdminUseCase {
	return &adminUseCase{
		adminRepo:  adminRepo,
		mailConfig: mailConfig,
	}
}
