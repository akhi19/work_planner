package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/akhi19/work_planner/configs"
	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/repository"
	"github.com/akhi19/work_planner/pkg/shift"
	"github.com/akhi19/work_planner/pkg/worker"
)

func main() {
	common.InitializeConnection(configs.GetConfig())

	shiftContaier := repository.ShiftContainer{}
	shiftContaier.Build()

	workerShitContainer := repository.WorkerShiftContainer{}
	workerShitContainer.Build()

	workerContainer := repository.WorkerContainer{}
	workerContainer.Build()
	router := mux.NewRouter().StrictSlash(false)

	plannerRoute := router.PathPrefix(
		"/planner/v1/",
	).Subrouter()

	worker.NewHttpServer(
		plannerRoute,
		workerContainer,
		workerShitContainer,
		shiftContaier,
	)

	shift.NewHttpServer(
		plannerRoute,
		shiftContaier,
		workerShitContainer,
	)

	common.GetLogger().Info("Starting server with ports : " + configs.GetConfig().Port)

	http.ListenAndServe(":"+configs.GetConfig().Port, router)
}
