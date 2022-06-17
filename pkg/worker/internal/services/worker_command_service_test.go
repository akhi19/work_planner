package services

import (
	"context"
	"errors"
	"testing"

	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/repository"
	mock_ports "github.com/akhi19/work_planner/pkg/repository/mocks"
	"github.com/akhi19/work_planner/pkg/worker/internal"
	"github.com/akhi19/work_planner/pkg/worker/internal/adaptors"
	"github.com/golang/mock/gomock"
)

var commandService WorkerCommandService
var mockWorkerContainer *mock_ports.MockIWorker
var mockShiftContainer *mock_ports.MockIShift
var mockWorkerShiftContainer *mock_ports.MockIWorkerShift

func InitCommandServiceTest(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockWorkerContainer = mock_ports.NewMockIWorker(ctrl)
	mockShiftContainer = mock_ports.NewMockIShift(ctrl)
	mockWorkerShiftContainer = mock_ports.NewMockIWorkerShift(ctrl)
	workerContainer := repository.WorkerContainer{
		IWorker: mockWorkerContainer,
	}
	shiftContainer := repository.ShiftContainer{
		IShift: mockShiftContainer,
	}
	workerShiftContainer := repository.WorkerShiftContainer{
		IWorkerShift: mockWorkerShiftContainer,
	}

	commandService = WorkerCommandService{
		repositoryAdaptor: adaptors.NewRepositoryAdaptor(
			workerContainer,
			workerShiftContainer,
			shiftContainer,
		),
	}
}

func TestWorkerCommandService_AddWorkerShift(t *testing.T) {
	InitCommandServiceTest(t)
	type args struct {
		ctx                   context.Context
		addWorkerShiftRequest internal.AddWorkerShiftRequestDTO
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		mockFunc func(args args)
	}{
		// VALID TEST CASE
		{
			name: "Valid addition",
			args: args{
				ctx: context.Background(),
				addWorkerShiftRequest: internal.AddWorkerShiftRequestDTO{
					WorkerID: 1,
					Date:     1,
					ShiftID:  1,
					Status:   "active",
				},
			},
			wantErr: false,
			mockFunc: func(args args) {
				mockShiftContainer.EXPECT().GetShiftByID(args.ctx, args.addWorkerShiftRequest.ShiftID).Return(
					&domain.ShiftDTO{
						ID: 1,
					},
					nil,
				).Times(1)

				mockWorkerContainer.EXPECT().GetWorkerByID(args.ctx, args.addWorkerShiftRequest.WorkerID).Return(
					&domain.WorkerDTO{
						ID: 1,
					},
					nil,
				).Times(1)

				mockWorkerShiftContainer.EXPECT().GetWorkerFromShift(args.ctx, args.addWorkerShiftRequest.WorkerID, args.addWorkerShiftRequest.Date).Return(
					nil,
					nil,
				).Times(1)

				workerShiftDTO := args.addWorkerShiftRequest.ToWorkerShiftDTO()
				mockWorkerShiftContainer.EXPECT().Insert(args.ctx, workerShiftDTO).Return(
					nil,
				).Times(1)
			},
		},

		// SHIFT NOT FOUND TEST CASE
		{
			name: "Invalid Shift",
			args: args{
				ctx: context.Background(),
				addWorkerShiftRequest: internal.AddWorkerShiftRequestDTO{
					WorkerID: 1,
					Date:     1,
					ShiftID:  1,
					Status:   "active",
				},
			},
			wantErr: true,
			mockFunc: func(args args) {
				mockShiftContainer.EXPECT().GetShiftByID(args.ctx, args.addWorkerShiftRequest.ShiftID).Return(
					nil,
					nil,
				).Times(1)
			},
		},

		// WORKER NOT FOUND TEST CASE
		{
			name: "Invalid worker",
			args: args{
				ctx: context.Background(),
				addWorkerShiftRequest: internal.AddWorkerShiftRequestDTO{
					WorkerID: 1,
					Date:     1,
					ShiftID:  1,
					Status:   "active",
				},
			},
			wantErr: true,
			mockFunc: func(args args) {
				mockShiftContainer.EXPECT().GetShiftByID(args.ctx, args.addWorkerShiftRequest.ShiftID).Return(
					&domain.ShiftDTO{
						ID: 1,
					},
					nil,
				).Times(1)

				mockWorkerContainer.EXPECT().GetWorkerByID(args.ctx, args.addWorkerShiftRequest.WorkerID).Return(
					nil,
					nil,
				).Times(1)
			},
		},

		// WORKER ASSIGNED FOR DAY TEST CASE
		{
			name: "Worker already assigned work",
			args: args{
				ctx: context.Background(),
				addWorkerShiftRequest: internal.AddWorkerShiftRequestDTO{
					WorkerID: 1,
					Date:     1,
					ShiftID:  1,
					Status:   "active",
				},
			},
			wantErr: true,
			mockFunc: func(args args) {
				mockShiftContainer.EXPECT().GetShiftByID(args.ctx, args.addWorkerShiftRequest.ShiftID).Return(
					&domain.ShiftDTO{
						ID: 1,
					},
					nil,
				).Times(1)

				mockWorkerContainer.EXPECT().GetWorkerByID(args.ctx, args.addWorkerShiftRequest.WorkerID).Return(
					&domain.WorkerDTO{
						ID: 1,
					},
					nil,
				).Times(1)

				id := domain.SqlID(1)
				mockWorkerShiftContainer.EXPECT().GetWorkerFromShift(args.ctx, args.addWorkerShiftRequest.WorkerID, args.addWorkerShiftRequest.Date).Return(
					&id,
					nil,
				).Times(1)
			},
		},

		// INSERT WORKER SHIFT ERROR TEST CASE
		{
			name: "Insert worker shift error test case",
			args: args{
				ctx: context.Background(),
				addWorkerShiftRequest: internal.AddWorkerShiftRequestDTO{
					WorkerID: 1,
					Date:     1,
					ShiftID:  1,
					Status:   "active",
				},
			},
			wantErr: true,
			mockFunc: func(args args) {
				mockShiftContainer.EXPECT().GetShiftByID(args.ctx, args.addWorkerShiftRequest.ShiftID).Return(
					&domain.ShiftDTO{
						ID: 1,
					},
					nil,
				).Times(1)

				mockWorkerContainer.EXPECT().GetWorkerByID(args.ctx, args.addWorkerShiftRequest.WorkerID).Return(
					&domain.WorkerDTO{
						ID: 1,
					},
					nil,
				).Times(1)

				mockWorkerShiftContainer.EXPECT().GetWorkerFromShift(args.ctx, args.addWorkerShiftRequest.WorkerID, args.addWorkerShiftRequest.Date).Return(
					nil,
					nil,
				).Times(1)

				workerShiftDTO := args.addWorkerShiftRequest.ToWorkerShiftDTO()
				mockWorkerShiftContainer.EXPECT().Insert(args.ctx, workerShiftDTO).Return(
					errors.New(""),
				).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.args)
			if err := commandService.AddWorkerShift(tt.args.ctx, tt.args.addWorkerShiftRequest); (err != nil) != tt.wantErr {
				t.Errorf("WorkerCommandService.AddWorkerShift() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
