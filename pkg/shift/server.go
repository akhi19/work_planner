package shift

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/repository"
	"github.com/akhi19/work_planner/pkg/shift/internal/adaptors"
	"github.com/akhi19/work_planner/pkg/shift/internal/ports"
)

func NewHttpServer(
	router *mux.Router,
	shiftContainer repository.ShiftContainer,
	workerShiftContainer repository.WorkerShiftContainer,
) {
	repositoryAdaptor := adaptors.NewRepositoryAdaptor(
		shiftContainer,
		workerShiftContainer,
	)

	clientPort := ports.NewClientPort(
		repositoryAdaptor,
	)

	router.HandleFunc(
		"/shifts",
		common.HttpRequestHandler(
			clientPort.AddShift,
		),
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/shifts/{id}",
		common.HttpRequestHandler(
			clientPort.DeleteShift,
		),
	).Methods(http.MethodDelete)

	router.HandleFunc(
		"/shifts",
		common.HttpRequestHandler(
			clientPort.GetShifts,
		),
	).Methods(http.MethodGet)
}
