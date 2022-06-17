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
	shiftDetailsTableName = "shift_details"
)

type ShiftAdaptor struct {
	sqlHandler *sql.DB
}

func NewShiftAdaptor() *ShiftAdaptor {
	return &ShiftAdaptor{
		sqlHandler: common.GetSqlHandler(),
	}
}

func (adaptor *ShiftAdaptor) Insert(
	ctx context.Context,
	shiftModel models.ShiftModel,
) error {
	queryStatement := fmt.Sprintf(`INSERT INTO %s(from_time, to_time, status) VALUES (%v, %v, '%s');`,
		shiftDetailsTableName,
		shiftModel.FromTimestamp,
		shiftModel.ToTimestamp,
		shiftModel.Status,
	)

	_, err := adaptor.sqlHandler.ExecContext(ctx,
		queryStatement,
	)
	if err != nil {
		return err
	}
	return nil
}

func (adaptor *ShiftAdaptor) Delete(
	ctx context.Context,
	id domain.SqlID,
) error {
	queryStatement := fmt.Sprintf(`UPDATE %s SET status = '%s' WHERE id = %v;`,
		shiftDetailsTableName,
		domain.EntityStatusInactive,
		id,
	)

	_, err := adaptor.sqlHandler.ExecContext(ctx,
		queryStatement,
	)
	return err
}

func (adaptor *ShiftAdaptor) GetShiftDetails(
	ctx context.Context,
) ([]models.ShiftModel, error) {
	queryStatement := fmt.Sprintf(`SELECT id, from_time, to_time, status FROM %s WHERE status='%s' ORDER BY from_time ASC;`,
		shiftDetailsTableName,
		domain.EntityStatusActive,
	)

	query, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
	)
	if err != nil {
		return nil, err
	}
	shifts := make([]models.ShiftModel, 0)
	defer query.Close()
	for query.Next() {
		shift := models.ShiftModel{}
		err = query.Scan(&shift.ID, &shift.FromTimestamp, &shift.ToTimestamp, &shift.Status)
		if err != nil {
			return nil, err
		}
		shifts = append(shifts, shift)
	}
	return shifts, nil
}

func (adaptor *ShiftAdaptor) GetShiftByID(
	ctx context.Context,
	shiftID domain.SqlID,
) (*models.ShiftModel, error) {
	queryStatement := fmt.Sprintf(`SELECT id, from_time, to_time, status FROM %s WHERE id = %v AND status = '%s';`,
		shiftDetailsTableName,
		shiftID,
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
	shift := models.ShiftModel{}
	err = query.Scan(&shift.ID, &shift.FromTimestamp, &shift.ToTimestamp, &shift.Status)
	if err != nil {
		return nil, err
	}
	return &shift, nil
}
