package ports

import (
	"context"

	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/worker/internal"
	"github.com/akhi19/work_planner/pkg/worker/internal/adaptors"
	"github.com/akhi19/work_planner/pkg/worker/internal/services"
)

var workerCommandService *services.WorkerCommandService
var workerQueryService *services.WorkerQueryService

type iWorkerCommandService interface {
	AddWorker(
		ctx context.Context,
		addWorkerRequest internal.AddWorkerRequestDTO,
	) error

	UpdateWorker(
		ctx context.Context,
		id domain.SqlID,
		updateWorkerRequest internal.UpdateWorkerRequestDTO,
	) error

	DeleteWorker(
		ctx context.Context,
		id domain.SqlID,
	) error

	AddWorkerShift(
		ctx context.Context,
		addWorkerShiftRequest internal.AddWorkerShiftRequestDTO,
	) error

	UpdateWorkerShift(
		ctx context.Context,
		id domain.SqlID,
		updateWorkerShiftRequest internal.UpdateWorkerShiftRequestDTO,
	) error

	DeleteWorkerShift(
		ctx context.Context,
		id domain.SqlID,
	) error
}

type iWorkerQueryService interface {
	GetWorkers(
		ctx context.Context,
	) ([]domain.WorkerDTO, error)

	GetFreeWorkers(
		ctx context.Context,
		date int64,
	) ([]domain.WorkerDTO, error)

	GetWorkersOccupied(
		ctx context.Context,
		date int64,
	) ([]domain.WorkerDTO, error)
}

func getWorkerCommandService(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) *services.WorkerCommandService {
	if workerCommandService == nil {
		workerCommandService = services.NewWorkerCommandService(
			repositoryAdaptor,
		)
	}
	return workerCommandService
}

func getWorkerQueryService(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) *services.WorkerQueryService {
	if workerQueryService == nil {
		workerQueryService = services.NewWorkerQueryService(
			repositoryAdaptor,
		)
	}
	return workerQueryService
}
