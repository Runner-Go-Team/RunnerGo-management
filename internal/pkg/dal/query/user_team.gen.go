// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"kp-management/internal/pkg/dal/model"
)

func newUserTeam(db *gorm.DB) userTeam {
	_userTeam := userTeam{}

	_userTeam.userTeamDo.UseDB(db)
	_userTeam.userTeamDo.UseModel(&model.UserTeam{})

	tableName := _userTeam.userTeamDo.TableName()
	_userTeam.ALL = field.NewAsterisk(tableName)
	_userTeam.ID = field.NewInt64(tableName, "id")
	_userTeam.UserID = field.NewInt64(tableName, "user_id")
	_userTeam.UserIdentify = field.NewString(tableName, "user_identify")
	_userTeam.TeamID = field.NewInt64(tableName, "team_id")
	_userTeam.RoleID = field.NewInt64(tableName, "role_id")
	_userTeam.InviteUserID = field.NewInt64(tableName, "invite_user_id")
	_userTeam.InviteUserIdentify = field.NewString(tableName, "invite_user_identify")
	_userTeam.Sort = field.NewInt32(tableName, "sort")
	_userTeam.CreatedAt = field.NewTime(tableName, "created_at")
	_userTeam.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userTeam.DeletedAt = field.NewField(tableName, "deleted_at")

	_userTeam.fillFieldMap()

	return _userTeam
}

type userTeam struct {
	userTeamDo userTeamDo

	ALL                field.Asterisk
	ID                 field.Int64
	UserID             field.Int64
	UserIdentify       field.String
	TeamID             field.Int64
	RoleID             field.Int64
	InviteUserID       field.Int64
	InviteUserIdentify field.String
	Sort               field.Int32
	CreatedAt          field.Time
	UpdatedAt          field.Time
	DeletedAt          field.Field

	fieldMap map[string]field.Expr
}

func (u userTeam) Table(newTableName string) *userTeam {
	u.userTeamDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userTeam) As(alias string) *userTeam {
	u.userTeamDo.DO = *(u.userTeamDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userTeam) updateTableName(table string) *userTeam {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UserID = field.NewInt64(table, "user_id")
	u.UserIdentify = field.NewString(table, "user_identify")
	u.TeamID = field.NewInt64(table, "team_id")
	u.RoleID = field.NewInt64(table, "role_id")
	u.InviteUserID = field.NewInt64(table, "invite_user_id")
	u.InviteUserIdentify = field.NewString(table, "invite_user_identify")
	u.Sort = field.NewInt32(table, "sort")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userTeam) WithContext(ctx context.Context) *userTeamDo { return u.userTeamDo.WithContext(ctx) }

func (u userTeam) TableName() string { return u.userTeamDo.TableName() }

func (u userTeam) Alias() string { return u.userTeamDo.Alias() }

func (u *userTeam) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userTeam) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 11)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["user_identify"] = u.UserIdentify
	u.fieldMap["team_id"] = u.TeamID
	u.fieldMap["role_id"] = u.RoleID
	u.fieldMap["invite_user_id"] = u.InviteUserID
	u.fieldMap["invite_user_identify"] = u.InviteUserIdentify
	u.fieldMap["sort"] = u.Sort
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userTeam) clone(db *gorm.DB) userTeam {
	u.userTeamDo.ReplaceDB(db)
	return u
}

type userTeamDo struct{ gen.DO }

func (u userTeamDo) Debug() *userTeamDo {
	return u.withDO(u.DO.Debug())
}

func (u userTeamDo) WithContext(ctx context.Context) *userTeamDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userTeamDo) ReadDB() *userTeamDo {
	return u.Clauses(dbresolver.Read)
}

func (u userTeamDo) WriteDB() *userTeamDo {
	return u.Clauses(dbresolver.Write)
}

func (u userTeamDo) Clauses(conds ...clause.Expression) *userTeamDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userTeamDo) Returning(value interface{}, columns ...string) *userTeamDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userTeamDo) Not(conds ...gen.Condition) *userTeamDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userTeamDo) Or(conds ...gen.Condition) *userTeamDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userTeamDo) Select(conds ...field.Expr) *userTeamDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userTeamDo) Where(conds ...gen.Condition) *userTeamDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userTeamDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userTeamDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userTeamDo) Order(conds ...field.Expr) *userTeamDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userTeamDo) Distinct(cols ...field.Expr) *userTeamDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userTeamDo) Omit(cols ...field.Expr) *userTeamDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userTeamDo) Join(table schema.Tabler, on ...field.Expr) *userTeamDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userTeamDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userTeamDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userTeamDo) RightJoin(table schema.Tabler, on ...field.Expr) *userTeamDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userTeamDo) Group(cols ...field.Expr) *userTeamDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userTeamDo) Having(conds ...gen.Condition) *userTeamDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userTeamDo) Limit(limit int) *userTeamDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userTeamDo) Offset(offset int) *userTeamDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userTeamDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userTeamDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userTeamDo) Unscoped() *userTeamDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userTeamDo) Create(values ...*model.UserTeam) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userTeamDo) CreateInBatches(values []*model.UserTeam, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userTeamDo) Save(values ...*model.UserTeam) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userTeamDo) First() (*model.UserTeam, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTeam), nil
	}
}

func (u userTeamDo) Take() (*model.UserTeam, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTeam), nil
	}
}

func (u userTeamDo) Last() (*model.UserTeam, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTeam), nil
	}
}

func (u userTeamDo) Find() ([]*model.UserTeam, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserTeam), err
}

func (u userTeamDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserTeam, err error) {
	buf := make([]*model.UserTeam, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userTeamDo) FindInBatches(result *[]*model.UserTeam, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userTeamDo) Attrs(attrs ...field.AssignExpr) *userTeamDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userTeamDo) Assign(attrs ...field.AssignExpr) *userTeamDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userTeamDo) Joins(fields ...field.RelationField) *userTeamDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userTeamDo) Preload(fields ...field.RelationField) *userTeamDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userTeamDo) FirstOrInit() (*model.UserTeam, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTeam), nil
	}
}

func (u userTeamDo) FirstOrCreate() (*model.UserTeam, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserTeam), nil
	}
}

func (u userTeamDo) FindByPage(offset int, limit int) (result []*model.UserTeam, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userTeamDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userTeamDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userTeamDo) Delete(models ...*model.UserTeam) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userTeamDo) withDO(do gen.Dao) *userTeamDo {
	u.DO = *do.(*gen.DO)
	return u
}
