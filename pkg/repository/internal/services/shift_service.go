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
) (*domain.SqlID, error) {

	shiftModel := models.ShiftModel{}
	shiftModel.FromShiftDTO(shiftDTO)

	id, err := service.adaptor.Insert(
		ctx,
		shiftModel,
	)
	if err != nil {
		return nil, err
	}
	sqlID := domain.SqlID(*id)
	return &sqlID, nil
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

func (service *ShiftService) GetShift(
	ctx context.Context,
	date int64,
) ([]domain.ShiftDTO, error) {
	shifts, err := service.adaptor.GetShifts(
		ctx,
		date,
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