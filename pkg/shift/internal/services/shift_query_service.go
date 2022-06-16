package services

import (
	"context"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/shift/internal/adaptors"
	"github.com/sirupsen/logrus"
)

type ShiftQueryService struct {
	repositoryAdaptor *adaptors.RepositoryAdaptor
}

func NewShiftQueryService(
	repository *adaptors.RepositoryAdaptor,
) *ShiftQueryService {
	return &ShiftQueryService{
		repositoryAdaptor: repository,
	}
}

func (service *ShiftQueryService) GetShifts(
	ctx context.Context,
	date int64,
) ([]domain.ShiftDTO, error) {
	log := common.GetLogger().WithFields(logrus.Fields{"function": "GetShifts"})
	shifts, err := service.repositoryAdaptor.ShiftContainer().IShift.GetShifts(
		ctx,
		date,
	)
	if err != nil {
		log.Error(err.Error())
		return nil, common.InternalServerError()
	}
	return shifts, nil
}
