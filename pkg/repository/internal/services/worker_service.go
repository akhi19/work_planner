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
) (*domain.SqlID, error) {

	workerModel := models.WorkerModel{}
	workerModel.FromWorkerDTO(workerDTO)

	id, err := service.adaptor.Insert(
		ctx,
		workerModel,
	)
	if err != nil {
		return nil, err
	}
	sqlID := domain.SqlID(*id)
	return &sqlID, nil
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
