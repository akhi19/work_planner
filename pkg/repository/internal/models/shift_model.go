package models

import (
	"github.com/akhi19/work_planner/pkg/domain"
)

type ShiftModel struct {
	ID            domain.SqlID
	FromTimestamp int64
	ToTimestamp   int64
	Status        domain.EntityStatus
}

func (entity *ShiftModel) FromShiftDTO(shiftDTO domain.ShiftDTO) {
	entity.FromTimestamp = shiftDTO.FromTimestamp
	entity.ToTimestamp = shiftDTO.ToTimestamp
	entity.Status = shiftDTO.Status
}

func (entity *ShiftModel) ToShiftDTO() domain.ShiftDTO {
	return domain.ShiftDTO{
		ID:            entity.ID,
		FromTimestamp: entity.FromTimestamp,
		ToTimestamp:   entity.ToTimestamp,
		Status:        entity.Status,
	}
}
