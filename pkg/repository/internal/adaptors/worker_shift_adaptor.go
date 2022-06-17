package adaptors

import (
	"context"
	"database/sql"
	"fmt"

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
) error {
	queryStatement := fmt.Sprintf(`INSERT INTO %s(worker_id, shift_id, date, status) VALUES (%v, %v, '%v', '%s');`,
		workerShiftDetailsTableName,
		shiftModel.WorkerID,
		shiftModel.ShiftID,
		common.TimeFromMillis(shiftModel.Date).Format("2006-01-02"),
		shiftModel.Status,
	)

	query, err := adaptor.sqlHandler.QueryContext(
		ctx,
		queryStatement,
	)
	if err != nil {
		return err
	}
	defer query.Close()
	return nil
}

func (adaptor *WorkerShiftAdaptor) Update(
	ctx context.Context,
	id domain.SqlID,
	updateModel models.UpdateWorkerShiftModel,
) error {
	queryStatement := fmt.Sprintf(`UPDATE %s SET shift_id = %v WHERE id = %v;`,
		workerShiftDetailsTableName,
		updateModel.ShiftID,
		id,
	)

	_, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
	)
	return err
}

func (adaptor *WorkerShiftAdaptor) Delete(
	ctx context.Context,
	id domain.SqlID,
) error {
	queryStatement := fmt.Sprintf(`UPDATE %s SET status = '%s' WHERE id = %v;`,
		workerShiftDetailsTableName,
		domain.EntityStatusInactive,
		id,
	)

	_, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
	)
	return err
}

func (adaptor *WorkerShiftAdaptor) DeleteUsingWorkerID(
	ctx context.Context,
	workerID domain.SqlID,
) error {
	queryStatement := fmt.Sprintf(`UPDATE %s SET status = '%s' WHERE worker_id = %v;`,
		workerShiftDetailsTableName,
		domain.EntityStatusInactive,
		workerID,
	)

	_, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
	)
	return err
}

func (adaptor *WorkerShiftAdaptor) DeleteUsingShiftID(
	ctx context.Context,
	shiftID domain.SqlID,
) error {
	queryStatement := fmt.Sprintf(`UPDATE %s SET status = '%s' WHERE shift_id = %v;`,
		workerShiftDetailsTableName,
		domain.EntityStatusInactive,
		shiftID,
	)

	_, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
	)
	return err
}

func (adaptor *WorkerShiftAdaptor) GetWorkersOccupied(
	ctx context.Context,
	date int64,
) ([]models.WorkerOccupiedModel, error) {
	queryStatement := fmt.Sprintf(`SELECT t1.id, t1.shift_id, t2.id, t2.name, t2.email, t2.phone, t2.status FROM %s AS t1 
	INNER JOIN
	%s AS t2
	ON t1.date = '%v'
	AND t1.status = '%s' AND t2.Status = '%s' AND t1.worker_id = t2.id;`,
		workerShiftDetailsTableName,
		workerDetailsTableName,
		common.TimeFromMillis(date).Format("2006-01-02"),
		domain.EntityStatusActive,
		domain.EntityStatusActive,
	)

	query, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
	)
	if err != nil {
		return nil, err
	}
	workerOccupiedDetails := make([]models.WorkerOccupiedModel, 0)
	defer query.Close()
	for query.Next() {
		workerOccupiedDetail := models.WorkerOccupiedModel{}
		err = query.Scan(&workerOccupiedDetail.ID, &workerOccupiedDetail.ShiftID, &workerOccupiedDetail.WorkerID, &workerOccupiedDetail.Name,
			&workerOccupiedDetail.Email, &workerOccupiedDetail.Phone, &workerOccupiedDetail.Status)
		if err != nil {
			return nil, err
		}
		workerOccupiedDetails = append(workerOccupiedDetails, workerOccupiedDetail)
	}
	return workerOccupiedDetails, nil
}

func (adaptor *WorkerShiftAdaptor) GetFreeWorkers(
	ctx context.Context,
	date int64,
) ([]models.WorkerModel, error) {
	queryStatement := fmt.Sprintf(`
    SELECT t1.id, t1.name, t1.email, t1.phone, t1.status FROM %s AS t1 
	LEFT JOIN
	%s AS t2
	ON t1.id = t2.worker_id AND t2.date IN ('%v', NULL) AND t2.status = '%s'
	WHERE t2.id IS NULL AND t1.status = '%s';`,
		workerDetailsTableName,
		workerShiftDetailsTableName,
		common.TimeFromMillis(date).Format("2006-01-02"),
		domain.EntityStatusActive,
		domain.EntityStatusActive,
	)

	query, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
	)
	if err != nil {
		return nil, err
	}
	workerDetails := make([]models.WorkerModel, 0)
	defer query.Close()
	for query.Next() {
		workerDetail := models.WorkerModel{}
		err = query.Scan(&workerDetail.ID, &workerDetail.Name, &workerDetail.Email, &workerDetail.Phone, &workerDetail.Status)
		if err != nil {
			return nil, err
		}
		workerDetails = append(workerDetails, workerDetail)
	}
	return workerDetails, nil
}

func (adaptor *WorkerShiftAdaptor) GetWorkerFromShift(
	ctx context.Context,
	workerID domain.SqlID,
	date int64,
) (*domain.SqlID, error) {
	queryStatement := fmt.Sprintf(`SELECT id FROM %s WHERE date = '%v' AND worker_id = %v AND status = '%s';`,
		workerShiftDetailsTableName,
		common.TimeFromMillis(date).Format("2006-01-02"),
		workerID,
		domain.EntityStatusActive,
	)

	query, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
	)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	if !query.Next() {
		return nil, nil
	}
	var id domain.SqlID
	err = query.Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
