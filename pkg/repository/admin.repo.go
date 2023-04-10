package repository

import (
	"database/sql"

	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
)

type adminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) interfaces.AdminRepository {
	return &adminRepo{
		db: db,
	}
}
