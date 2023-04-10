package repository

import (
	"database/sql"

	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
)

type workerRepository struct {
	db *sql.DB
}

func NewWorkerRepo(db *sql.DB) interfaces.WorkerRepository {
	return &workerRepository{
		db: db,
	}
}
