package models

import (
	"github.com/akhi19/work_planner/pkg/domain"
)

type ShiftModel struct {
	ID       domain.SqlID
	FromTime int64
	ToTime   int64
	Status   domain.EntityStatus
}

func (entity *ShiftModel) FromShiftDTO(shiftDTO domain.ShiftDTO) {
	entity.FromTime = shiftDTO.FromTime
	entity.ToTime = shiftDTO.ToTime
	entity.Status = shiftDTO.Status
}

func (entity *ShiftModel) ToShiftDTO() domain.ShiftDTO {
	return domain.ShiftDTO{
		ID:       entity.ID,
		FromTime: entity.FromTime,
		ToTime:   entity.ToTime,
		Status:   entity.Status,
	}
}
