package ports

import (
	"context"

	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/shift/internal"
	"github.com/akhi19/work_planner/pkg/shift/internal/adaptors"
	"github.com/akhi19/work_planner/pkg/shift/internal/services"
)

var shiftCommandService *services.ShiftCommandService
var shiftQueryService *services.ShiftQueryService

type iShiftCommandService interface {
	AddShift(
		ctx context.Context,
		addShiftRequestDTO internal.AddShiftRequestDTO,
	) error

	DeleteShift(
		ctx context.Context,
		shiftID domain.SqlID,
	) error
}

type iShiftQueryService interface {
	GetShifts(
		ctx context.Context,
		date int64,
	) ([]domain.ShiftDTO, error)
}

func getShiftCommandService(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) *services.ShiftCommandService {
	if shiftCommandService == nil {
		shiftCommandService = services.NewShiftCommandService(
			repositoryAdaptor,
		)
	}
	return shiftCommandService
}

func getShiftQueryService(
	repositoryAdaptor *adaptors.RepositoryAdaptor,
) *services.ShiftQueryService {
	if shiftQueryService == nil {
		shiftQueryService = services.NewShiftQueryService(
			repositoryAdaptor,
		)
	}
	return shiftQueryService
}
