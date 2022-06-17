package ports

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/worker/internal"
	"github.com/akhi19/work_planner/pkg/worker/internal/adaptors"
)

type ClientPort struct {
	iWorkerCommandService iWorkerCommandService
	iWorkerQueryService   iWorkerQueryService
}

func NewClientPort(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) ClientPort {
	return ClientPort{
		iWorkerCommandService: getWorkerCommandService(
			repositoryAdaptor,
		),
		iWorkerQueryService: getWorkerQueryService(
			repositoryAdaptor,
		),
	}
}

func (port *ClientPort) AddWorker(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	var addWorkerRequestDTO internal.AddWorkerRequestDTO
	err := addWorkerRequestDTO.Populate(
		request.Body,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	err = port.iWorkerCommandService.AddWorker(
		request.Context(),
		addWorkerRequestDTO,
	)

	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendEmptyHttpResponse(
		responseWriter,
		http.StatusOK,
	)
}

func (port *ClientPort) UpdateWorker(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	var updateWorkerRequestDTO internal.UpdateWorkerRequestDTO
	err := updateWorkerRequestDTO.Populate(
		request.Body,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	urlParams := mux.Vars(request)
	sqlIDString := urlParams["id"]

	sqlID, err := strconv.Atoi(sqlIDString)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	err = port.iWorkerCommandService.UpdateWorker(
		request.Context(),
		domain.SqlID(sqlID),
		updateWorkerRequestDTO,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendEmptyHttpResponse(
		responseWriter,
		http.StatusOK,
	)
}

func (port *ClientPort) DeleteWorker(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	urlParams := mux.Vars(request)
	sqlIDString := urlParams["id"]

	sqlID, err := strconv.Atoi(sqlIDString)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	err = port.iWorkerCommandService.DeleteWorker(
		request.Context(),
		domain.SqlID(sqlID),
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendEmptyHttpResponse(
		responseWriter,
		http.StatusOK,
	)
}

func (port *ClientPort) AddWorkerShift(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	queryParams := request.URL.Query()
	dateString := queryParams.Get("date")

	time, err := common.GetParsedDate(dateString, time.UTC.String())
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}
	var addWorkerShiftRequestDTO internal.AddWorkerShiftRequestDTO
	err = addWorkerShiftRequestDTO.Populate(
		request.Body,
		*time,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	err = port.iWorkerCommandService.AddWorkerShift(
		request.Context(),
		addWorkerShiftRequestDTO,
	)

	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendEmptyHttpResponse(
		responseWriter,
		http.StatusOK,
	)
}

func (port *ClientPort) UpdateWorkerShift(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	var updateWorkerShiftRequestDTO internal.UpdateWorkerShiftRequestDTO
	err := updateWorkerShiftRequestDTO.Populate(
		request.Body,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	urlParams := mux.Vars(request)
	sqlIDString := urlParams["id"]

	sqlID, err := strconv.Atoi(sqlIDString)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	err = port.iWorkerCommandService.UpdateWorkerShift(
		request.Context(),
		domain.SqlID(sqlID),
		updateWorkerShiftRequestDTO,
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendEmptyHttpResponse(
		responseWriter,
		http.StatusOK,
	)
}

func (port *ClientPort) DeleteWorkerShift(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	urlParams := mux.Vars(request)
	sqlIDString := urlParams["id"]

	sqlID, err := strconv.Atoi(sqlIDString)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	err = port.iWorkerCommandService.DeleteWorkerShift(
		request.Context(),
		domain.SqlID(sqlID),
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendEmptyHttpResponse(
		responseWriter,
		http.StatusOK,
	)
}

func (port *ClientPort) GetWorkers(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	workers, err := port.iWorkerQueryService.GetWorkers(
		request.Context(),
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendHttpResponse(
		responseWriter,
		http.StatusOK,
		workers,
	)
}

func (port *ClientPort) GetWorkersOccupied(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	queryParams := request.URL.Query()
	dateString := queryParams.Get("date")

	time, err := common.GetParsedDate(dateString, time.UTC.String())
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	workers, err := port.iWorkerQueryService.GetWorkersOccupied(
		request.Context(),
		time.UnixMilli(),
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendHttpResponse(
		responseWriter,
		http.StatusOK,
		workers,
	)
}

func (port *ClientPort) GetFreeWorkers(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	queryParams := request.URL.Query()
	dateString := queryParams.Get("date")

	time, err := common.GetParsedDate(dateString, time.UTC.String())
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	workers, err := port.iWorkerQueryService.GetFreeWorkers(
		request.Context(),
		time.UnixMilli(),
	)
	if err != nil {
		common.SendHttpError(
			request.Context(),
			responseWriter,
			err,
		)
		return
	}

	common.SendHttpResponse(
		responseWriter,
		http.StatusOK,
		workers,
	)
}
