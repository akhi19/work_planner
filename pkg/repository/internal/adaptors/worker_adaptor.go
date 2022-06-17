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
) error {
	queryStatement := fmt.Sprintf(`INSERT INTO %s(name, email, phone, status ) VALUES ('%s', '%s', %v, '%s');`,
		workerDetailsTableName,
		workerModel.Name,
		workerModel.Email,
		workerModel.Phone,
		workerModel.Status,
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

func (adaptor *WorkerAdaptor) Update(
	ctx context.Context,
	id domain.SqlID,
	updateModel models.UpdateWorkerModel,
) error {
	queryStatement := fmt.Sprintf(`UPDATE %s SET name = '%s', phone = '%s' WHERE id = %v;`,
		workerDetailsTableName,
		updateModel.Name,
		updateModel.Phone,
		id,
	)

	_, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
	)
	return err
}

func (adaptor *WorkerAdaptor) Delete(
	ctx context.Context,
	id domain.SqlID,
) error {
	queryStatement := fmt.Sprintf(`UPDATE %s SET status = '%s' WHERE id = %v;`,
		workerDetailsTableName,
		domain.EntityStatusInactive,
		id,
	)

	_, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
	)
	return err
}

func (adaptor *WorkerAdaptor) GetWorkers(
	ctx context.Context,
) ([]models.WorkerModel, error) {
	queryStatement := fmt.Sprintf(`SELECT id, name, email, phone, status FROM %s WHERE status='%s' ORDER BY name ASC;`,
		workerDetailsTableName,
		domain.EntityStatusActive,
	)

	query, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
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

func (adaptor *WorkerAdaptor) GetWorkerByID(
	ctx context.Context,
	workerID domain.SqlID,
) (*models.WorkerModel, error) {
	queryStatement := fmt.Sprintf(`SELECT id, name, email, phone, status FROM %s WHERE id = %v AND status = '%s';`,
		workerDetailsTableName,
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
	worker := models.WorkerModel{}
	err = query.Scan(&worker.ID, &worker.Name, &worker.Email, &worker.Phone, &worker.Status)
	if err != nil {
		return nil, err
	}
	return &worker, nil
}
