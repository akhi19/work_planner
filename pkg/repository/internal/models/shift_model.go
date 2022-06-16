package models

import (
	"github.com/akhi19/work_planner/pkg/domain"
)

type ShiftModel struct {
	ID            domain.SqlID
	DateTimestamp int64
	FromTimestamp int64
	ToTimestamp   int64
	Status        domain.EntityStatus
}

func (entity *ShiftModel) FromShiftDTO(shiftDTO domain.ShiftDTO) {
	entity.DateTimestamp = shiftDTO.DateTimestamp
	entity.FromTimestamp = shiftDTO.FromTimestamp
	entity.ToTimestamp = shiftDTO.ToTimestamp
	entity.Status = shiftDTO.Status
}

func (entity *ShiftModel) ToShiftDTO() domain.ShiftDTO {
	return domain.ShiftDTO{
		ID:            entity.ID,
		DateTimestamp: entity.DateTimestamp,
		FromTimestamp: entity.FromTimestamp,
		ToTimestamp:   entity.ToTimestamp,
		Status:        entity.Status,
	}
}
