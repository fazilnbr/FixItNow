package handler

import (
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type WorkerHandler struct {
	workerService services.WorkerUseCase
}

func NewWorkerHandler(workerService services.WorkerUseCase) WorkerHandler {
	return WorkerHandler{
		workerService: workerService,
	}
}
