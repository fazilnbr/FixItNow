package usecase

import (
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type workerService struct {
	workerRepo interfaces.WorkerRepository
}

func NewWorkerService(workerRepo interfaces.WorkerRepository) services.WorkerUseCase {
	return &workerService{
		workerRepo: workerRepo,
	}

}
