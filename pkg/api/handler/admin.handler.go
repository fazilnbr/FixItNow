package handler

import (
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type AdminHandler struct {
	adminService services.AdminUseCase
}

func NewAdminHandler(adminService services.AdminUseCase) AdminHandler {
	return AdminHandler{
		adminService: adminService,
	}
}
