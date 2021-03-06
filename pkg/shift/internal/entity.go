package internal

import (
	"encoding/json"
	"io"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
)

type AddShiftRequestDTO struct {
	FromTime *int64 `json:"from_time" validate:"required,min=0,max=24"`
	ToTime   *int64 `json:"to_time" validate:"required,min=0,max=24"`
}

func (entity *AddShiftRequestDTO) Populate(
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

func (entity *AddShiftRequestDTO) ToShiftDTO() domain.ShiftDTO {
	return domain.ShiftDTO{
		FromTime: *entity.FromTime,
		ToTime:   *entity.ToTime,
		Status:   domain.EntityStatusActive,
	}
}
