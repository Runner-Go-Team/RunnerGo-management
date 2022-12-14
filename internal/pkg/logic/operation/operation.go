package operation

import (
	"context"
	"time"

	"github.com/go-omnibus/omnibus"

	"kp-management/internal/pkg/dal"
	"kp-management/internal/pkg/dal/query"
	"kp-management/internal/pkg/dal/rao"
	"kp-management/internal/pkg/packer"
)

func List(ctx context.Context, teamID int64, limit, offset int) ([]*rao.Operation, int64, error) {
	tx := query.Use(dal.DB()).Operation
	operations, cnt, err := tx.WithContext(ctx).Where(tx.TeamID.Eq(teamID), tx.CreatedAt.Between(time.Now().AddDate(0, 0, -2), time.Now())).
		Order(tx.UpdatedAt.Desc()).FindByPage(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	var userIDs []int64
	for _, o := range operations {
		userIDs = append(userIDs, o.UserID)
	}

	u := query.Use(dal.DB()).User
	users, err := u.WithContext(ctx).Where(u.ID.In(omnibus.Int64ArrayUnique(userIDs)...)).Find()
	if err != nil {
		return nil, 0, err
	}

	return packer.TransOperationsToRaoOperationList(operations, users), cnt, nil
}
