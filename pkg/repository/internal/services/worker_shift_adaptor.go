package services

import (
	"context"

	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/repository/internal/adaptors"
	"github.com/akhi19/work_planner/pkg/repository/internal/models"
)

type WorkerShiftService struct {
	adaptor *adaptors.WorkerShiftAdaptor
}

func NewWorkerShiftService(
	adaptor *adaptors.WorkerShiftAdaptor,
) *WorkerShiftService {
	return &WorkerShiftService{
		adaptor: adaptor,
	}
}

func (service *WorkerShiftService) Insert(
	ctx context.Context,
	workerShiftDTO domain.WorkerShiftDTO,
) (*domain.SqlID, error) {

	workerShiftModel := models.WorkerShiftModel{}
	workerShiftModel.FromWorkerShiftDTO(workerShiftDTO)

	id, err := service.adaptor.Insert(
		ctx,
		workerShiftModel,
	)
	if err != nil {
		return nil, err
	}
	sqlID := domain.SqlID(*id)
	return &sqlID, nil
}

func (service *WorkerShiftService) Update(
	ctx context.Context,
	id domain.SqlID,
	updateWorkerShiftDTO domain.UpdateWorkerShiftDTO,
) error {

	updateWorkerShiftModel := models.UpdateWorkerShiftModel{}
	updateWorkerShiftModel.FromUpdateWorkerShiftDTO(updateWorkerShiftDTO)

	return service.adaptor.Update(
		ctx,
		id,
		updateWorkerShiftModel,
	)
}

func (service *WorkerShiftService) Delete(
	ctx context.Context,
	id domain.SqlID,
) error {

	return service.adaptor.Delete(
		ctx,
		id,
	)
}

func (service *WorkerShiftService) DeleteUsingWorkerID(
	ctx context.Context,
	workerID domain.SqlID,
) error {
	return service.adaptor.DeleteUsingWorkerID(
		ctx,
		workerID,
	)
}

func (service *WorkerShiftService) DeleteUsingShiftID(
	ctx context.Context,
	shiftID domain.SqlID,
) error {
	return service.adaptor.DeleteUsingWorkerID(
		ctx,
		shiftID,
	)
}

func (service *WorkerShiftService) GetFreeWorkers(
	ctx context.Context,
	date int64,
) ([]domain.WorkerDTO, error) {
	freeWorkers, err := service.adaptor.GetFreeWorkers(
		ctx,
		date,
	)
	if err != nil {
		return nil, err
	}
	workerDTOs := make([]domain.WorkerDTO, len(freeWorkers))
	for i, worker := range freeWorkers {
		workerDTOs[i] = worker.ToWorkerDTO()
	}
	return workerDTOs, nil
}

func (service *WorkerShiftService) GetWorkersOccupied(
	ctx context.Context,
	date int64,
) ([]domain.WorkerDTO, error) {
	freeWorkers, err := service.adaptor.GetWorkersOccupied(
		ctx,
		date,
	)
	if err != nil {
		return nil, err
	}
	workerDTOs := make([]domain.WorkerDTO, len(freeWorkers))
	for i, worker := range freeWorkers {
		workerDTOs[i] = worker.ToWorkerDTO()
	}
	return workerDTOs, nil
}

func (service *WorkerShiftService) GetWorkerFromShift(
	ctx context.Context,
	workerID domain.SqlID,
	date int64,
) (*domain.SqlID, error) {
	id, err := service.adaptor.GetWorkerFromShift(
		ctx,
		workerID,
		date,
	)
	if err != nil {
		return nil, err
	}
	return id, nil
}
