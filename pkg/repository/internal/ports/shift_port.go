package ports

import (
	"context"

	"github.com/akhi19/work_planner/pkg/domain"
)

type IShift interface {
	Insert(
		ctx context.Context,
		shiftDTO domain.ShiftDTO,
	) (*domain.SqlID, error)

	Delete(
		ctx context.Context,
		id domain.SqlID,
	) error

	GetShift(
		ctx context.Context,
		date int64,
	) ([]domain.ShiftDTO, error)
}
