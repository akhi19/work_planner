package worker

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/repository"
	"github.com/akhi19/work_planner/pkg/worker/internal/adaptors"
	"github.com/akhi19/work_planner/pkg/worker/internal/ports"
)

func NewHttpServer(
	router *mux.Router,
	workerContainer repository.WorkerContainer,
	workerShiftContainer repository.WorkerShiftContainer,
) {
	repositoryAdaptor := adaptors.NewRepositoryAdaptor(
		workerContainer,
		workerShiftContainer,
	)

	clientPort := ports.NewClientPort(
		repositoryAdaptor,
	)

	router.HandleFunc(
		"/workers",
		common.HttpRequestHandler(
			clientPort.AddWorker,
		),
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/workers/{id}",
		common.HttpRequestHandler(
			clientPort.UpdateWorker,
		),
	).Methods(http.MethodPut)

	router.HandleFunc(
		"/workers/{id}",
		common.HttpRequestHandler(
			clientPort.DeleteWorker,
		),
	).Methods(http.MethodDelete)

	router.HandleFunc(
		"/workers-shift",
		common.HttpRequestHandler(
			clientPort.AddWorkerShift,
		),
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/workers-shift/{id}",
		common.HttpRequestHandler(
			clientPort.UpdateWorkerShift,
		),
	).Methods(http.MethodPut)

	router.HandleFunc(
		"/workers-shift/{id}",
		common.HttpRequestHandler(
			clientPort.DeleteWorker,
		),
	).Methods(http.MethodDelete)

	router.HandleFunc(
		"/workers",
		common.HttpRequestHandler(
			clientPort.GetWorkers,
		),
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/workers/free",
		common.HttpRequestHandler(
			clientPort.GetFreeWorkers,
		),
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/workers/occupied",
		common.HttpRequestHandler(
			clientPort.GetWorkersOccupied,
		),
	).Methods(http.MethodGet)
}
