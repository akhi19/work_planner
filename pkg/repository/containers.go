package repository

import (
	"github.com/akhi19/work_planner/pkg/repository/internal/adaptors"
	"github.com/akhi19/work_planner/pkg/repository/internal/ports"
	"github.com/akhi19/work_planner/pkg/repository/internal/services"
)

type ShiftContainer struct {
	IShift ports.IShift
}

type WorkerShiftContainer struct {
	IWorkerShift ports.IWorkerShift
}

type WorkerContainer struct {
	IWorker ports.IWorker
}

func (container *ShiftContainer) Build() {
	shiftAdaptor := adaptors.NewShiftAdaptor()
	container.IShift = services.NewShiftService(
		shiftAdaptor,
	)
}

func (container *WorkerShiftContainer) Build() {
	workerShiftAdaptor := adaptors.NewWorkerShiftAdaptor()
	container.IWorkerShift = services.NewWorkerShiftService(
		workerShiftAdaptor,
	)
}

func (container *WorkerContainer) Build() {
	workerAdaptor := adaptors.NewWorkerAdaptor()
	container.IWorker = services.NewWorkerService(
		workerAdaptor,
	)
}
