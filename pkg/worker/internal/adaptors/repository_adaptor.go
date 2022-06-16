package adaptors

import (
	"github.com/akhi19/work_planner/pkg/repository"
)

type RepositoryAdaptor struct {
	workerContainer      repository.WorkerContainer
	workerShiftContainer repository.WorkerShiftContainer
	shiftContainer       repository.ShiftContainer
}

func NewRepositoryAdaptor(
	workerContainer repository.WorkerContainer,
	workerShiftContainer repository.WorkerShiftContainer,
	shiftContainer repository.ShiftContainer,
) *RepositoryAdaptor {
	return &RepositoryAdaptor{
		workerContainer:      workerContainer,
		workerShiftContainer: workerShiftContainer,
		shiftContainer:       shiftContainer,
	}
}

func (adaptor *RepositoryAdaptor) WorkerContainer() repository.WorkerContainer {
	return adaptor.workerContainer
}

func (adaptor *RepositoryAdaptor) WorkerShiftContainer() repository.WorkerShiftContainer {
	return adaptor.workerShiftContainer
}

func (adaptor *RepositoryAdaptor) ShiftContainer() repository.ShiftContainer {
	return adaptor.shiftContainer
}
