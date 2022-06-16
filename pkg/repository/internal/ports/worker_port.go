package ports

import (
	"context"

	"github.com/akhi19/work_planner/pkg/domain"
)

type IWorker interface {
	Insert(
		ctx context.Context,
		workerDTO domain.WorkerDTO,
	) (*domain.SqlID, error)

	Update(
		ctx context.Context,
		id domain.SqlID,
		updateWorkerDTO domain.UpdateWorkerDTO,
	) error

	Delete(
		ctx context.Context,
		id domain.SqlID,
	) error

	GetWorkers(
		ctx context.Context,
	) ([]domain.WorkerDTO, error)
}
