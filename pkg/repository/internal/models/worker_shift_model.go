package models

import (
	"github.com/akhi19/work_planner/pkg/domain"
)

type WorkerShiftModel struct {
	ID       domain.SqlID
	WorkerID domain.SqlID
	ShiftID  domain.SqlID
	Date     int64
	Status   domain.EntityStatus
}

func (entity *WorkerShiftModel) FromWorkerShiftDTO(workerShiftDTO domain.WorkerShiftDTO) {
	entity.WorkerID = workerShiftDTO.WorkerID
	entity.ShiftID = workerShiftDTO.ShiftID
	entity.Status = workerShiftDTO.Status
	entity.Date = workerShiftDTO.Date
}

func (entity *WorkerShiftModel) ToWorkerShiftDTO() domain.WorkerShiftDTO {
	return domain.WorkerShiftDTO{
		ID:       entity.ID,
		WorkerID: entity.WorkerID,
		ShiftID:  entity.ShiftID,
		Date:     entity.Date,
		Status:   entity.Status,
	}
}

type UpdateWorkerShiftModel struct {
	ShiftID domain.SqlID
	Date    int64
}

func (entity *UpdateWorkerShiftModel) FromUpdateWorkerShiftDTO(workerShiftDTO domain.UpdateWorkerShiftDTO) {
	entity.ShiftID = workerShiftDTO.ShiftID
	entity.Date = workerShiftDTO.Date
}
