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

func newOperation(db *gorm.DB) operation {
	_operation := operation{}

	_operation.operationDo.UseDB(db)
	_operation.operationDo.UseModel(&model.Operation{})

	tableName := _operation.operationDo.TableName()
	_operation.ALL = field.NewAsterisk(tableName)
	_operation.ID = field.NewInt64(tableName, "id")
	_operation.TeamID = field.NewInt64(tableName, "team_id")
	_operation.UserID = field.NewInt64(tableName, "user_id")
	_operation.Category = field.NewInt32(tableName, "category")
	_operation.Name = field.NewString(tableName, "name")
	_operation.CreatedAt = field.NewTime(tableName, "created_at")
	_operation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_operation.DeletedAt = field.NewField(tableName, "deleted_at")

	_operation.fillFieldMap()

	return _operation
}

type operation struct {
	operationDo operationDo

	ALL       field.Asterisk
	ID        field.Int64
	TeamID    field.Int64
	UserID    field.Int64
	Category  field.Int32
	Name      field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (o operation) Table(newTableName string) *operation {
	o.operationDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o operation) As(alias string) *operation {
	o.operationDo.DO = *(o.operationDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *operation) updateTableName(table string) *operation {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.TeamID = field.NewInt64(table, "team_id")
	o.UserID = field.NewInt64(table, "user_id")
	o.Category = field.NewInt32(table, "category")
	o.Name = field.NewString(table, "name")
	o.CreatedAt = field.NewTime(table, "created_at")
	o.UpdatedAt = field.NewTime(table, "updated_at")
	o.DeletedAt = field.NewField(table, "deleted_at")

	o.fillFieldMap()

	return o
}

func (o *operation) WithContext(ctx context.Context) *operationDo {
	return o.operationDo.WithContext(ctx)
}

func (o operation) TableName() string { return o.operationDo.TableName() }

func (o operation) Alias() string { return o.operationDo.Alias() }

func (o *operation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *operation) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 8)
	o.fieldMap["id"] = o.ID
	o.fieldMap["team_id"] = o.TeamID
	o.fieldMap["user_id"] = o.UserID
	o.fieldMap["category"] = o.Category
	o.fieldMap["name"] = o.Name
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["updated_at"] = o.UpdatedAt
	o.fieldMap["deleted_at"] = o.DeletedAt
}

func (o operation) clone(db *gorm.DB) operation {
	o.operationDo.ReplaceDB(db)
	return o
}

type operationDo struct{ gen.DO }

func (o operationDo) Debug() *operationDo {
	return o.withDO(o.DO.Debug())
}

func (o operationDo) WithContext(ctx context.Context) *operationDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o operationDo) ReadDB() *operationDo {
	return o.Clauses(dbresolver.Read)
}

func (o operationDo) WriteDB() *operationDo {
	return o.Clauses(dbresolver.Write)
}

func (o operationDo) Clauses(conds ...clause.Expression) *operationDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o operationDo) Returning(value interface{}, columns ...string) *operationDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o operationDo) Not(conds ...gen.Condition) *operationDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o operationDo) Or(conds ...gen.Condition) *operationDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o operationDo) Select(conds ...field.Expr) *operationDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o operationDo) Where(conds ...gen.Condition) *operationDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o operationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *operationDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o operationDo) Order(conds ...field.Expr) *operationDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o operationDo) Distinct(cols ...field.Expr) *operationDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o operationDo) Omit(cols ...field.Expr) *operationDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o operationDo) Join(table schema.Tabler, on ...field.Expr) *operationDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o operationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *operationDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o operationDo) RightJoin(table schema.Tabler, on ...field.Expr) *operationDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o operationDo) Group(cols ...field.Expr) *operationDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o operationDo) Having(conds ...gen.Condition) *operationDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o operationDo) Limit(limit int) *operationDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o operationDo) Offset(offset int) *operationDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o operationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *operationDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o operationDo) Unscoped() *operationDo {
	return o.withDO(o.DO.Unscoped())
}

func (o operationDo) Create(values ...*model.Operation) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o operationDo) CreateInBatches(values []*model.Operation, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o operationDo) Save(values ...*model.Operation) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o operationDo) First() (*model.Operation, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Operation), nil
	}
}

func (o operationDo) Take() (*model.Operation, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Operation), nil
	}
}

func (o operationDo) Last() (*model.Operation, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Operation), nil
	}
}

func (o operationDo) Find() ([]*model.Operation, error) {
	result, err := o.DO.Find()
	return result.([]*model.Operation), err
}

func (o operationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Operation, err error) {
	buf := make([]*model.Operation, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o operationDo) FindInBatches(result *[]*model.Operation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o operationDo) Attrs(attrs ...field.AssignExpr) *operationDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o operationDo) Assign(attrs ...field.AssignExpr) *operationDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o operationDo) Joins(fields ...field.RelationField) *operationDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o operationDo) Preload(fields ...field.RelationField) *operationDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o operationDo) FirstOrInit() (*model.Operation, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Operation), nil
	}
}

func (o operationDo) FirstOrCreate() (*model.Operation, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Operation), nil
	}
}

func (o operationDo) FindByPage(offset int, limit int) (result []*model.Operation, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o operationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o operationDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o operationDo) Delete(models ...*model.Operation) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *operationDo) withDO(do gen.Dao) *operationDo {
	o.DO = *do.(*gen.DO)
	return o
}
