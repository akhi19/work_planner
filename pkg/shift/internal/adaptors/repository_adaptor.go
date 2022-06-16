package adaptors

import (
	"github.com/akhi19/work_planner/pkg/repository"
)

type RepositoryAdaptor struct {
	shiftContainer       repository.ShiftContainer
	workerShiftContainer repository.WorkerShiftContainer
}

func NewRepositoryAdaptor(
	shiftContainer repository.ShiftContainer,
	workerShiftContainer repository.WorkerShiftContainer,
) *RepositoryAdaptor {
	return &RepositoryAdaptor{
		shiftContainer:       shiftContainer,
		workerShiftContainer: workerShiftContainer,
	}
}

func (adaptor *RepositoryAdaptor) ShiftContainer() repository.ShiftContainer {
	return adaptor.shiftContainer
}

func (adaptor *RepositoryAdaptor) WorkerShiftContainer() repository.WorkerShiftContainer {
	return adaptor.workerShiftContainer
}
