package services

import (
	"context"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/worker/internal/adaptors"
	"github.com/sirupsen/logrus"
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
	log := common.GetLogger().WithFields(logrus.Fields{"function": "GetWorkers"})
	workers, err := service.repositoryAdaptor.WorkerContainer().IWorker.GetWorkers(
		ctx,
	)
	if err != nil {
		log.Error(err.Error())
		return nil, common.InternalServerError()
	}
	return workers, nil
}

func (service *WorkerQueryService) GetFreeWorkers(
	ctx context.Context,
	date int64,
) ([]domain.WorkerDTO, error) {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "GetFreeWorkers"})
	workers, err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.GetFreeWorkers(
		ctx,
		date,
	)
	if err != nil {
		log.Error(err.Error())
		return nil, common.InternalServerError()
	}
	return workers, nil
}

func (service *WorkerQueryService) GetWorkersOccupied(
	ctx context.Context,
	date int64,
) ([]domain.WorkerDTO, error) {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "GetWorkersOccupied"})
	workers, err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.GetWorkersOccupied(
		ctx,
		date,
	)
	if err != nil {
		log.Error(err.Error())
		return nil, common.InternalServerError()
	}
	return workers, nil
}
