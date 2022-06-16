package adaptors

import (
	"context"
	"database/sql"

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
) (*int64, error) {
	queryStatement := `
    INSERT INTO @Table(from_timestamp, to_timestamp) VALUES ( @FromTimestamp, @ToTimestamp);
    select isNull(SCOPE_IDENTITY(), -1);
   `

	query, err := adaptor.sqlHandler.Prepare(queryStatement)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	newRecord := query.QueryRowContext(ctx,
		sql.Named("Table", shiftDetailsTableName),
		sql.Named("FromTimestamp", shiftModel.FromTimestamp),
		sql.Named("ToTimestamp", shiftModel.ToTimestamp),
	)

	var newID int64
	err = newRecord.Scan(&newID)
	if err != nil {
		return nil, err
	}

	return &newID, nil
}

func (adaptor *ShiftAdaptor) Delete(
	ctx context.Context,
	id domain.SqlID,
) error {
	queryStatement := `
    UPDATE @Table SET status = @Status WHERE id = @ID;
   `

	_, err := adaptor.sqlHandler.ExecContext(ctx,
		queryStatement,
		sql.Named("Table", shiftDetailsTableName),
		sql.Named("Status", domain.EntityStatusInactive),
		sql.Named("ID", id),
	)
	return err
}

func (adaptor *ShiftAdaptor) GetShifts(
	ctx context.Context,
	date int64,
) ([]models.ShiftModel, error) {
	queryStatement := `
    SELECT id, from_timestamp, to_timestamp, status FROM @Table WHERE date = @Date AND status = @Status ORDER BY from_timestamp ASC;
   `

	query, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
		sql.Named("Table", shiftDetailsTableName),
		sql.Named("Date", date),
		sql.Named("Status", domain.EntityStatusActive),
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
	queryStatement := `
    SELECT id, from_timestamp, to_timestamp, status FROM @Table WHERE shift_id = @ShiftID;
   `

	query, err := adaptor.sqlHandler.QueryContext(ctx,
		queryStatement,
		sql.Named("Table", shiftDetailsTableName),
		sql.Named("ShiftID", shiftID),
		sql.Named("Status", domain.EntityStatusActive),
	)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	shift := models.ShiftModel{}
	err = query.Scan(&shift.ID, &shift.FromTimestamp, &shift.ToTimestamp, &shift.Status)
	if err != nil {
		return nil, err
	}
	return &shift, nil
}
