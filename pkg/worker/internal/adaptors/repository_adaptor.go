package adaptors

import (
	"github.com/akhi19/work_planner/pkg/repository"
)

type RepositoryAdaptor struct {
	workerContainer      repository.WorkerContainer
	workerShiftContainer repository.WorkerShiftContainer
}

func NewRepositoryAdaptor(
	workerContainer repository.WorkerContainer,
	workerShiftContainer repository.WorkerShiftContainer,
) *RepositoryAdaptor {
	return &RepositoryAdaptor{
		workerContainer:      workerContainer,
		workerShiftContainer: workerShiftContainer,
	}
}

func (adaptor *RepositoryAdaptor) WorkerContainer() repository.WorkerContainer {
	return adaptor.workerContainer
}

func (adaptor *RepositoryAdaptor) WorkerShiftContainer() repository.WorkerShiftContainer {
	return adaptor.workerShiftContainer
}
