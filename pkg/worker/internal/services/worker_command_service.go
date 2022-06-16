package services

import (
	"context"
	"fmt"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/worker/internal"
	"github.com/akhi19/work_planner/pkg/worker/internal/adaptors"
)

type WorkerCommandService struct {
	repositoryAdaptor *adaptors.RepositoryAdaptor
}

func NewWorkerCommandService(
	repository *adaptors.RepositoryAdaptor,
) *WorkerCommandService {
	return &WorkerCommandService{
		repositoryAdaptor: repository,
	}
}

func (service *WorkerCommandService) AddWorker(
	ctx context.Context,
	addWorkerRequest internal.AddWorkerRequestDTO,
) error {
	workerDTO := addWorkerRequest.ToWorkerDTO()
	_, err := service.repositoryAdaptor.WorkerContainer().IWorker.Insert(
		ctx,
		workerDTO,
	)
	if err != nil {
		//TODO : Add logs
		fmt.Println(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *WorkerCommandService) UpdateWorker(
	ctx context.Context,
	id domain.SqlID,
	updateWorkerRequest internal.UpdateWorkerRequestDTO,
) error {
	updateWorkerDTO := updateWorkerRequest.ToUpdateWorkerDTO()
	err := service.repositoryAdaptor.WorkerContainer().IWorker.Update(
		ctx,
		id,
		updateWorkerDTO,
	)
	if err != nil {
		//TODO : Add logs
		fmt.Println(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *WorkerCommandService) DeleteWorker(
	ctx context.Context,
	id domain.SqlID,
) error {
	err := service.repositoryAdaptor.WorkerContainer().IWorker.Delete(
		ctx,
		id,
	)
	if err != nil {
		//TODO : Add logs
		fmt.Println(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *WorkerCommandService) AddWorkerShift(
	ctx context.Context,
	addWorkerShiftRequest internal.AddWorkerShiftRequestDTO,
) error {
	workerShiftDTO := addWorkerShiftRequest.ToWorkerShiftDTO()
	_, err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.Insert(
		ctx,
		workerShiftDTO,
	)
	if err != nil {
		//TODO : Add logs
		fmt.Println(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *WorkerCommandService) UpdateWorkerShift(
	ctx context.Context,
	id domain.SqlID,
	updateWorkerShiftRequest internal.UpdateWorkerShiftRequestDTO,
) error {
	updateWorkerShiftDTO := updateWorkerShiftRequest.ToUpdateWorkerShiftDTO()
	err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.Update(
		ctx,
		id,
		updateWorkerShiftDTO,
	)
	if err != nil {
		//TODO : Add logs
		fmt.Println(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *WorkerCommandService) DeleteWorkerShift(
	ctx context.Context,
	id domain.SqlID,
) error {
	err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.Delete(
		ctx,
		id,
	)
	if err != nil {
		//TODO : Add logs
		fmt.Println(err.Error())
		return common.InternalServerError()
	}
	return nil
}
