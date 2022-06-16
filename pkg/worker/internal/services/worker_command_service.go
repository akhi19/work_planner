package services

import (
	"context"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/worker/internal"
	"github.com/akhi19/work_planner/pkg/worker/internal/adaptors"
	"github.com/sirupsen/logrus"
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
	log := common.GetLogger().WithFields(logrus.Fields{"function": "AddWorker"})
	workerDTO := addWorkerRequest.ToWorkerDTO()
	_, err := service.repositoryAdaptor.WorkerContainer().IWorker.Insert(
		ctx,
		workerDTO,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *WorkerCommandService) UpdateWorker(
	ctx context.Context,
	id domain.SqlID,
	updateWorkerRequest internal.UpdateWorkerRequestDTO,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "UpdateWorker"})
	updateWorkerDTO := updateWorkerRequest.ToUpdateWorkerDTO()
	err := service.repositoryAdaptor.WorkerContainer().IWorker.Update(
		ctx,
		id,
		updateWorkerDTO,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *WorkerCommandService) DeleteWorker(
	ctx context.Context,
	id domain.SqlID,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "UpdateWorker"})
	err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.DeleteUsingWorkerID(
		ctx,
		id,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}

	err = service.repositoryAdaptor.WorkerContainer().IWorker.Delete(
		ctx,
		id,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *WorkerCommandService) AddWorkerShift(
	ctx context.Context,
	addWorkerShiftRequest internal.AddWorkerShiftRequestDTO,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "UpdateWorker"})
	workerShiftDTO := addWorkerShiftRequest.ToWorkerShiftDTO()

	shiftDTO, err := service.repositoryAdaptor.ShiftContainer().IShift.GetShiftByID(
		ctx,
		addWorkerShiftRequest.ShiftID,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	if shiftDTO == nil {
		log.Error("no shift found")
		return common.BadRequest(
			common.BadRequestCode,
			"No shift found",
		)
	}

	id, err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.GetWorkerFromShift(
		ctx,
		addWorkerShiftRequest.WorkerID,
		addWorkerShiftRequest.Date,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	if id != nil {
		log.Error("worker assigned for day")
		return common.BadRequest(
			common.BadRequestCode,
			"Worker assigned for day",
		)
	}

	_, err = service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.Insert(
		ctx,
		workerShiftDTO,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *WorkerCommandService) UpdateWorkerShift(
	ctx context.Context,
	id domain.SqlID,
	updateWorkerShiftRequest internal.UpdateWorkerShiftRequestDTO,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "UpdateWorkerShift"})
	updateWorkerShiftDTO := updateWorkerShiftRequest.ToUpdateWorkerShiftDTO()
	shiftDTO, err := service.repositoryAdaptor.ShiftContainer().IShift.GetShiftByID(
		ctx,
		updateWorkerShiftDTO.ShiftID,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	if shiftDTO == nil {
		log.Error("no shift found")
		return common.BadRequest(
			common.BadRequestCode,
			"No shift found",
		)
	}
	err = service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.Update(
		ctx,
		id,
		updateWorkerShiftDTO,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *WorkerCommandService) DeleteWorkerShift(
	ctx context.Context,
	id domain.SqlID,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "UpdateWorker"})
	err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.Delete(
		ctx,
		id,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}
