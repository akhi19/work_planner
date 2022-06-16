package adaptors

import (
	"context"
	"database/sql"

	"github.com/akhi19/work_planner/pkg/common"
	"github.com/akhi19/work_planner/pkg/domain"
	"github.com/akhi19/work_planner/pkg/repository/internal/models"
)

const (
	workerDetailsTableName = "worker_details"
)

type WorkerAdaptor struct {
	sqlHandler *sql.DB
}

func NewWorkerAdaptor() *WorkerAdaptor {
	return &WorkerAdaptor{
		sqlHandler: common.GetSqlHandler(),
	}
}

func (adaptor *WorkerAdaptor) Insert(
	ctx context.Context,
	workerModel models.WorkerModel,
) (*int64, error) {
	queryStatement := `
    INSERT INTO @Table(name, email, phone, status ) VALUES (@Name, @Email, @Phone, @Status);
    select isNull(SCOPE_IDENTITY(), -1);
   `

	query, err := adaptor.sqlHandler.Prepare(queryStatement)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	newRecord := query.QueryRowContext(ctx,
		sql.Named("Table", workerDetailsTableName),
		sql.Named("Name", workerModel.Name),
		sql.Named("Email", workerModel.Email),
		sql.Named("Phone", workerModel.Phone),
		sql.Named("Status", workerModel.Status),
	)

	var newID int64
	err = newRecord.Scan(&newID)
	if err != nil {
		return nil, err
	}

	return &newID, nil
}

func (adaptor *WorkerAdaptor) Update(
	ctx context.Context,
	id domain.SqlID,
	updateModel models.UpdateWorkerModel,
) error {
	queryStatement := `
    UPDATE @Table SET name = @Name, phone = @Phone WHERE id = @ID;
   `

	_, err := adaptor.sqlHandler.ExecContext(ctx,
		queryStatement,
		sql.Named("Table", workerDetailsTableName),
		sql.Named("Name", updateModel.Name),
		sql.Named("Phone", updateModel.Phone),
		sql.Named("ID", id),
	)
	return err
}

func (adaptor *WorkerAdaptor) Delete(
	ctx context.Context,
	id domain.SqlID,
) error {
	queryStatement := `
	UPDATE @Table SET status = @Status WHERE id = @ID;
   `

	_, err := adaptor.sqlHandler.ExecContext(ctx,
		queryStatement,
		sql.Named("Table", workerDetailsTableName),
		sql.Named("Status", domain.EntityStatusInactive),
		sql.Named("ID", id),
	)
	return err
}

func (adaptor *WorkerAdaptor) GetWorkers(
	ctx context.Context,
) ([]models.WorkerModel, error) {
	queryStatement := `
    SELECT id, name, email, phone, status FROM @Table WHERE status = @Status;
   `
	query, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
		sql.Named("Table", workerDetailsTableName),
		sql.Named("Status", domain.EntityStatusActive),
	)
	if err != nil {
		return nil, err
	}
	workers := make([]models.WorkerModel, 0)
	defer query.Close()
	for query.Next() {
		worker := models.WorkerModel{}
		err = query.Scan(&worker.ID, &worker.Name, &worker.Email, &worker.Phone, &worker.Status)
		if err != nil {
			return nil, err
		}
		workers = append(workers, worker)
	}
	return workers, nil
}
