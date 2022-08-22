// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

func Use(db *gorm.DB) *Query {
	return &Query{
		db:        db,
		Group:     newGroup(db),
		Operation: newOperation(db),
		Plan:      newPlan(db),
		Report:    newReport(db),
		Target:    newTarget(db),
		Team:      newTeam(db),
		User:      newUser(db),
		UserTeam:  newUserTeam(db),
	}
}

type Query struct {
	db *gorm.DB

	Group     group
	Operation operation
	Plan      plan
	Report    report
	Target    target
	Team      team
	User      user
	UserTeam  userTeam
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:        db,
		Group:     q.Group.clone(db),
		Operation: q.Operation.clone(db),
		Plan:      q.Plan.clone(db),
		Report:    q.Report.clone(db),
		Target:    q.Target.clone(db),
		Team:      q.Team.clone(db),
		User:      q.User.clone(db),
		UserTeam:  q.UserTeam.clone(db),
	}
}

type queryCtx struct {
	Group     *groupDo
	Operation *operationDo
	Plan      *planDo
	Report    *reportDo
	Target    *targetDo
	Team      *teamDo
	User      *userDo
	UserTeam  *userTeamDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Group:     q.Group.WithContext(ctx),
		Operation: q.Operation.WithContext(ctx),
		Plan:      q.Plan.WithContext(ctx),
		Report:    q.Report.WithContext(ctx),
		Target:    q.Target.WithContext(ctx),
		Team:      q.Team.WithContext(ctx),
		User:      q.User.WithContext(ctx),
		UserTeam:  q.UserTeam.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
