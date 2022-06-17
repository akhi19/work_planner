package services

import (
	"context"

	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/repository/internal/adaptors"
	"github.com/akhi19/work_planner/pkg/repository/internal/models"
)

type ShiftService struct {
	adaptor *adaptors.ShiftAdaptor
}

func NewShiftService(
	adaptor *adaptors.ShiftAdaptor,
) *ShiftService {
	return &ShiftService{
		adaptor: adaptor,
	}
}

func (service *ShiftService) Insert(
	ctx context.Context,
	shiftDTO domain.ShiftDTO,
) error {

	shiftModel := models.ShiftModel{}
	shiftModel.FromShiftDTO(shiftDTO)

	err := service.adaptor.Insert(
		ctx,
		shiftModel,
	)
	if err != nil {
		return err
	}
	return nil
}

func (service *ShiftService) Delete(
	ctx context.Context,
	id domain.SqlID,
) error {

	return service.adaptor.Delete(
		ctx,
		id,
	)
}

func (service *ShiftService) GetShiftDetails(
	ctx context.Context,
) ([]domain.ShiftDTO, error) {
	shifts, err := service.adaptor.GetShiftDetails(
		ctx,
	)
	if err != nil {
		return nil, err
	}
	shiftDTOs := make([]domain.ShiftDTO, len(shifts))
	for i, shift := range shifts {
		shiftDTOs[i] = shift.ToShiftDTO()
	}
	return shiftDTOs, nil
}

func (service *ShiftService) GetShiftByID(
	ctx context.Context,
	shiftID domain.SqlID,
) (*domain.ShiftDTO, error) {
	shift, err := service.adaptor.GetShiftByID(
		ctx,
		shiftID,
	)
	if shift == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	shiftDTO := shift.ToShiftDTO()
	return &shiftDTO, nil
}
