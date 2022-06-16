package ports

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/shift/internal"
	"github.com/akhi19/work_planner/pkg/shift/internal/adaptors"
)

type ClientPort struct {
	iShiftCommandService iShiftCommandService
	iShiftQueryService   iShiftQueryService
}

func NewClientPort(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) ClientPort {
	return ClientPort{
		iShiftCommandService: getShiftCommandService(
			repositoryAdaptor,
		),
		iShiftQueryService: getShiftQueryService(
			repositoryAdaptor,
		),
	}
}

func (port *ClientPort) AddShift(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	var addShiftRequestDTO internal.AddShiftRequestDTO
	err := addShiftRequestDTO.Populate(
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

	err = port.iShiftCommandService.AddShift(
		request.Context(),
		addShiftRequestDTO,
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

func (port *ClientPort) DeleteShift(
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

	err = port.iShiftCommandService.DeleteShift(
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

func (port *ClientPort) GetShifts(
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

	workers, err := port.iShiftQueryService.GetShifts(
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
