package services

import (
	"context"
	"fmt"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/worker/internal/adaptors"
)

type WorkerQueryService struct {
	repositoryAdaptor *adaptors.RepositoryAdaptor
}

func NewWorkerQueryService(
	repository *adaptors.RepositoryAdaptor,
) *WorkerQueryService {
	return &WorkerQueryService{
		repositoryAdaptor: repository,
	}
}

func (service *WorkerQueryService) GetWorkers(
	ctx context.Context,
) ([]domain.WorkerDTO, error) {
	workers, err := service.repositoryAdaptor.WorkerContainer().IWorker.GetWorkers(
		ctx,
	)
	if err != nil {
		//TODO : Add logs
		fmt.Println(err.Error())
		return nil, common.InternalServerError()
	}
	return workers, nil
}

func (service *WorkerQueryService) GetFreeWorkers(
	ctx context.Context,
	date int64,
) ([]domain.WorkerDTO, error) {
	workers, err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.GetFreeWorkers(
		ctx,
		date,
	)
	if err != nil {
		//TODO : Add logs
		fmt.Println(err.Error())
		return nil, common.InternalServerError()
	}
	return workers, nil
}

func (service *WorkerQueryService) GetWorkersOccupied(
	ctx context.Context,
	date int64,
) ([]domain.WorkerDTO, error) {
	workers, err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.GetWorkersOccupied(
		ctx,
		date,
	)
	if err != nil {
		//TODO : Add logs
		fmt.Println(err.Error())
		return nil, common.InternalServerError()
	}
	return workers, nil
}
