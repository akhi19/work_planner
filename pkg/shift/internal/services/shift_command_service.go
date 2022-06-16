package services

import (
	"context"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/shift/internal"
	"github.com/akhi19/work_planner/pkg/shift/internal/adaptors"
	"github.com/sirupsen/logrus"
)

type ShiftCommandService struct {
	repositoryAdaptor *adaptors.RepositoryAdaptor
}

func NewShiftCommandService(
	repository *adaptors.RepositoryAdaptor,
) *ShiftCommandService {
	return &ShiftCommandService{
		repositoryAdaptor: repository,
	}
}

func (service *ShiftCommandService) AddShift(
	ctx context.Context,
	addShiftRequestDTO internal.AddShiftRequestDTO,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "AddShift"})
	shiftDTO := addShiftRequestDTO.ToShiftDTO()
	_, err := service.repositoryAdaptor.ShiftContainer().IShift.Insert(
		ctx,
		shiftDTO,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}

func (service *ShiftCommandService) DeleteShift(
	ctx context.Context,
	shiftID domain.SqlID,
) error {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "DeleteShift"})
	err := service.repositoryAdaptor.WorkerShiftContainer().IWorkerShift.DeleteUsingShiftID(
		ctx,
		shiftID,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}

	err = service.repositoryAdaptor.ShiftContainer().IShift.Delete(
		ctx,
		shiftID,
	)
	if err != nil {
		log.Error(err.Error())
		return common.InternalServerError()
	}
	return nil
}
