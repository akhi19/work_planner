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

	DeleteUsingWorkerID(
		ctx context.Context,
		workerID domain.SqlID,
	) error

	DeleteUsingShiftID(
		ctx context.Context,
		shiftID domain.SqlID,
	) error

	GetFreeWorkers(
		ctx context.Context,
		date int64,
	) ([]domain.WorkerDTO, error)

	GetWorkerFromShift(
		ctx context.Context,
		workerID domain.SqlID,
		date int64,
	) (*domain.SqlID, error)

	GetWorkersOccupied(
		ctx context.Context,
		date int64,
	) ([]domain.WorkerDTO, error)
}
