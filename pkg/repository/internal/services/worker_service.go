package services

import (
	"context"

	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/repository/internal/adaptors"
	"github.com/akhi19/work_planner/pkg/repository/internal/models"
)

type WorkerService struct {
	adaptor *adaptors.WorkerAdaptor
}

func NewWorkerService(
	adaptor *adaptors.WorkerAdaptor,
) *WorkerService {
	return &WorkerService{
		adaptor: adaptor,
	}
}

func (service *WorkerService) Insert(
	ctx context.Context,
	workerDTO domain.WorkerDTO,
) error {

	workerModel := models.WorkerModel{}
	workerModel.FromWorkerDTO(workerDTO)

	err := service.adaptor.Insert(
		ctx,
		workerModel,
	)
	if err != nil {
		return err
	}
	return nil
}

func (service *WorkerService) Update(
	ctx context.Context,
	id domain.SqlID,
	updateWorkerDTO domain.UpdateWorkerDTO,
) error {

	updateWorkerModel := models.UpdateWorkerModel{}
	updateWorkerModel.FromUpdateWorkerDTO(updateWorkerDTO)

	return service.adaptor.Update(
		ctx,
		id,
		updateWorkerModel,
	)
}

func (service *WorkerService) Delete(
	ctx context.Context,
	id domain.SqlID,
) error {

	return service.adaptor.Delete(
		ctx,
		id,
	)
}

func (service *WorkerService) GetWorkers(
	ctx context.Context,
) ([]domain.WorkerDTO, error) {
	workers, err := service.adaptor.GetWorkers(
		ctx,
	)
	if err != nil {
		return nil, err
	}
	workerDTOs := make([]domain.WorkerDTO, len(workers))
	for i, worker := range workers {
		workerDTOs[i] = worker.ToWorkerDTO()
	}
	return workerDTOs, nil
}

func (service *WorkerService) GetWorkerByID(
	ctx context.Context,
	workerID domain.SqlID,
) (*domain.WorkerDTO, error) {
	worker, err := service.adaptor.GetWorkerByID(
		ctx,
		workerID,
	)
	if worker == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	workerDTO := worker.ToWorkerDTO()
	return &workerDTO, nil
}
