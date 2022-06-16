package models

import (
	"github.com/akhi19/work_planner/pkg/domain"
)

type WorkerModel struct {
	ID     domain.SqlID
	Name   string
	Phone  string
	Email  string
	Status domain.EntityStatus
}

func (entity *WorkerModel) FromWorkerDTO(workerDTO domain.WorkerDTO) {
	entity.Name = workerDTO.Name
	entity.Phone = workerDTO.Phone
	entity.Email = workerDTO.Email
	entity.Status = workerDTO.Status
}

func (entity *WorkerModel) ToWorkerDTO() domain.WorkerDTO {
	return domain.WorkerDTO{
		ID:     entity.ID,
		Name:   entity.Name,
		Email:  entity.Email,
		Phone:  entity.Phone,
		Status: entity.Status,
	}
}

type UpdateWorkerModel struct {
	Name  string
	Phone string
}

func (entity *UpdateWorkerModel) FromUpdateWorkerDTO(workerDTO domain.UpdateWorkerDTO) {
	entity.Name = workerDTO.Name
	entity.Phone = workerDTO.Phone
}
