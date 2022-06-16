package adaptors

import (
	"context"
	"database/sql"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/repository/internal/models"
)

const (
	workerShiftDetailsTableName = "worker_shift_details"
)

type WorkerShiftAdaptor struct {
	sqlHandler *sql.DB
}

func NewWorkerShiftAdaptor() *WorkerShiftAdaptor {
	return &WorkerShiftAdaptor{
		sqlHandler: common.GetSqlHandler(),
	}
}

func (adaptor *WorkerShiftAdaptor) Insert(
	ctx context.Context,
	shiftModel models.WorkerShiftModel,
) (*int64, error) {
	queryStatement := `
    INSERT INTO @Table(worker_id, shift_id, status) VALUES (@WorkerID, @ShiftID, @Status);
    select isNull(SCOPE_IDENTITY(), -1);
   `

	query, err := adaptor.sqlHandler.Prepare(queryStatement)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	newRecord := query.QueryRowContext(ctx,
		sql.Named("Table", workerShiftDetailsTableName),
		sql.Named("WorkerID", shiftModel.WorkerID),
		sql.Named("ShiftID", shiftModel.ShiftID),
		sql.Named("Status", shiftModel.Status),
	)

	var newID int64
	err = newRecord.Scan(&newID)
	if err != nil {
		return nil, err
	}

	return &newID, nil
}

func (adaptor *WorkerShiftAdaptor) Update(
	ctx context.Context,
	id domain.SqlID,
	updateModel models.UpdateWorkerShiftModel,
) error {
	queryStatement := `
    UPDATE @Table SET shift_id = @ShiftID WHERE id = @ID;
   `

	_, err := adaptor.sqlHandler.ExecContext(ctx,
		queryStatement,
		sql.Named("Table", workerShiftDetailsTableName),
		sql.Named("ShiftID", updateModel.ShiftID),
		sql.Named("ID", id),
	)
	return err
}

func (adaptor *WorkerShiftAdaptor) Delete(
	ctx context.Context,
	id domain.SqlID,
) error {
	queryStatement := `
    UPDATE @Table SET status = @Status WHERE id = @ID;
   `

	_, err := adaptor.sqlHandler.ExecContext(ctx,
		queryStatement,
		sql.Named("Table", workerShiftDetailsTableName),
		sql.Named("Status", domain.EntityStatusInactive),
		sql.Named("ID", id),
	)
	return err
}

func (adaptor *WorkerShiftAdaptor) GetShiftWorkers(
	ctx context.Context,
	date int64,
) ([]models.WorkerModel, error) {
	queryStatement := `
    SELECT t2.id, t2.name, t2.email, t2.phone FROM @Table AS t1 
	INNER JOIN
	@WorkerTable AS t2
	ON t1.shift_id  IN (SELECT shift_id FROM @ShiftTable WHERE date = @Date AND status = @Status) 
	AND t1.status = @Status AND t2.Status = @Status AND t1.worker_id = t2.id;
   `

	query, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
		sql.Named("Table", shiftDetailsTableName),
		sql.Named("WorkerTable", workerDetailsTableName),
		sql.Named("ShiftTable", shiftDetailsTableName),
		sql.Named("Date", date),
		sql.Named("Status", domain.EntityStatusActive),
	)
	if err != nil {
		return nil, err
	}
	workerDetails := make([]models.WorkerModel, 0)
	defer query.Close()
	for query.Next() {
		workerDetail := models.WorkerModel{}
		err = query.Scan(&workerDetail.ID, workerDetail.Name, workerDetail.Email, workerDetail.Phone)
		if err != nil {
			return nil, err
		}
		workerDetails = append(workerDetails, workerDetail)
	}
	return workerDetails, nil
}

func (adaptor *WorkerShiftAdaptor) GetFreeWorkers(
	ctx context.Context,
	date int64,
) ([]models.WorkerModel, error) {
	queryStatement := `
    SELECT t2.id, t2.name, t2.email, t2.phone FROM @Table AS t1 
	RIGHT JOIN
	@WorkerTable AS t2
	ON t1.shift_id IN (SELECT shift_id FROM @ShiftTable WHERE date = @Date AND status = @Status) 
	AND t1.Status = @Status AND t2.Status = @Status AND t1.worker_id = t2.id
	WHERE t1.id IS NULL;
   `

	query, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
		sql.Named("Table", shiftDetailsTableName),
		sql.Named("WorkerTable", workerDetailsTableName),
		sql.Named("ShiftTable", shiftDetailsTableName),
		sql.Named("Date", date),
		sql.Named("Status", domain.EntityStatusActive),
	)
	if err != nil {
		return nil, err
	}
	workerDetails := make([]models.WorkerModel, 0)
	defer query.Close()
	for query.Next() {
		workerDetail := models.WorkerModel{}
		err = query.Scan(&workerDetail.ID, workerDetail.Name, workerDetail.Email, workerDetail.Phone)
		if err != nil {
			return nil, err
		}
		workerDetails = append(workerDetails, workerDetail)
	}
	return workerDetails, nil
}
