package internal

import (
	"encoding/json"
	"io"
	"regexp"
	"time"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type AddWorkerRequestDTO struct {
	Name  string `json:"name" validate:"required, max=200"`
	Email string `json:"email" validate:"required"`
	Phone int64  `json:"phone" validate:"required,min=10,max=10"`
}

func (entity *AddWorkerRequestDTO) Populate(
	body io.ReadCloser,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}

	err = common.Validator.Struct(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}
	if !emailRegex.MatchString(entity.Email) {
		return common.ValidationError(
			"Please provide valid mail",
		)
	}
	return nil
}

func (entity *AddWorkerRequestDTO) ToWorkerDTO() domain.WorkerDTO {
	return domain.WorkerDTO{
		Name:   entity.Name,
		Email:  entity.Email,
		Phone:  entity.Phone,
		Status: domain.EntityStatusActive,
	}
}

type UpdateWorkerRequestDTO struct {
	Name  string `json:"name" validate:"required,max=200"`
	Phone string `json:"phone" validate:"required,min=10,max=10"`
}

func (entity *UpdateWorkerRequestDTO) Populate(
	body io.ReadCloser,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}

	err = common.Validator.Struct(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}
	return nil
}

func (entity *UpdateWorkerRequestDTO) ToUpdateWorkerDTO() domain.UpdateWorkerDTO {
	return domain.UpdateWorkerDTO{
		Name:  entity.Name,
		Phone: entity.Phone,
	}
}

type AddWorkerShiftRequestDTO struct {
	WorkerID domain.SqlID        `json:"worker_id" validate:"required"`
	Date     int64               `json:"date"`
	ShiftID  domain.SqlID        `json:"shift_id" validate:"required"`
	Status   domain.EntityStatus `json:"status"`
}

func (entity *AddWorkerShiftRequestDTO) Populate(
	body io.ReadCloser,
	time time.Time,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}

	err = common.Validator.Struct(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}
	entity.Date = time.UnixMilli()
	entity.Status = domain.EntityStatusActive
	return nil
}

func (entity *AddWorkerShiftRequestDTO) ToWorkerShiftDTO() domain.WorkerShiftDTO {
	return domain.WorkerShiftDTO{
		WorkerID: entity.WorkerID,
		ShiftID:  entity.ShiftID,
		Status:   domain.EntityStatusActive,
		Date:     entity.Date,
	}
}

type UpdateWorkerShiftRequestDTO struct {
	ShiftID domain.SqlID `json:"shift_id" validate:"required"`
}

func (entity *UpdateWorkerShiftRequestDTO) Populate(
	body io.ReadCloser,
) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}

	err = common.Validator.Struct(entity)
	if err != nil {
		return common.ValidationError(
			common.InvalidRequestMsg,
		)
	}
	return nil
}

func (entity *UpdateWorkerShiftRequestDTO) ToUpdateWorkerShiftDTO() domain.UpdateWorkerShiftDTO {
	return domain.UpdateWorkerShiftDTO{
		ShiftID: entity.ShiftID,
	}
}
