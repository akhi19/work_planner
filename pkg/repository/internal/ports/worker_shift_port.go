package ports

import (
	"context"

	"github.com/akhi19/work_planner/pkg/domain"
)

type IWorkerShift interface {
	Insert(
		ctx context.Context,
		workerShiftDTO domain.WorkerShiftDTO,
	) (*domain.SqlID, error)

	Update(
		ctx context.Context,
		id domain.SqlID,
		updateWorkerShiftDTO domain.UpdateWorkerShiftDTO,
	) error

	Delete(
		ctx context.Context,
		id domain.SqlID,
	) error

	GetFreeWorkers(
		ctx context.Context,
		date int64,
	) ([]domain.WorkerDTO, error)

	GetShiftWorkers(
		ctx context.Context,
		date int64,
	) ([]domain.WorkerDTO, error)
}
