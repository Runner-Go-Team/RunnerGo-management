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

	"kp-management/cmd/generate/internal/pkg/dal/model"
)

func newMachine(db *gorm.DB, opts ...gen.DOOption) machine {
	_machine := machine{}

	_machine.machineDo.UseDB(db, opts...)
	_machine.machineDo.UseModel(&model.Machine{})

	tableName := _machine.machineDo.TableName()
	_machine.ALL = field.NewAsterisk(tableName)
	_machine.ID = field.NewInt64(tableName, "id")
	_machine.Region = field.NewString(tableName, "region")
	_machine.IP = field.NewString(tableName, "ip")
	_machine.Port = field.NewInt32(tableName, "port")
	_machine.Weight = field.NewInt32(tableName, "weight")
	_machine.Status = field.NewInt32(tableName, "status")
	_machine.CreatedAt = field.NewTime(tableName, "created_at")
	_machine.UpdatedAt = field.NewTime(tableName, "updated_at")
	_machine.DeletedAt = field.NewField(tableName, "deleted_at")

	_machine.fillFieldMap()

	return _machine
}

type machine struct {
	machineDo machineDo

	ALL       field.Asterisk
	ID        field.Int64  // id
	Region    field.String // 地域
	IP        field.String // ip
	Port      field.Int32  // 端口
	Weight    field.Int32  // 额外权重
	Status    field.Int32  // 机器状态{1: 空闲, 2: 忙碌}
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 修改时间
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (m machine) Table(newTableName string) *machine {
	m.machineDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m machine) As(alias string) *machine {
	m.machineDo.DO = *(m.machineDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *machine) updateTableName(table string) *machine {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.Region = field.NewString(table, "region")
	m.IP = field.NewString(table, "ip")
	m.Port = field.NewInt32(table, "port")
	m.Weight = field.NewInt32(table, "weight")
	m.Status = field.NewInt32(table, "status")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")

	m.fillFieldMap()

	return m
}

func (m *machine) WithContext(ctx context.Context) *machineDo { return m.machineDo.WithContext(ctx) }

func (m machine) TableName() string { return m.machineDo.TableName() }

func (m machine) Alias() string { return m.machineDo.Alias() }

func (m *machine) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *machine) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 9)
	m.fieldMap["id"] = m.ID
	m.fieldMap["region"] = m.Region
	m.fieldMap["ip"] = m.IP
	m.fieldMap["port"] = m.Port
	m.fieldMap["weight"] = m.Weight
	m.fieldMap["status"] = m.Status
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
}

func (m machine) clone(db *gorm.DB) machine {
	m.machineDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m machine) replaceDB(db *gorm.DB) machine {
	m.machineDo.ReplaceDB(db)
	return m
}

type machineDo struct{ gen.DO }

func (m machineDo) Debug() *machineDo {
	return m.withDO(m.DO.Debug())
}

func (m machineDo) WithContext(ctx context.Context) *machineDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m machineDo) ReadDB() *machineDo {
	return m.Clauses(dbresolver.Read)
}

func (m machineDo) WriteDB() *machineDo {
	return m.Clauses(dbresolver.Write)
}

func (m machineDo) Session(config *gorm.Session) *machineDo {
	return m.withDO(m.DO.Session(config))
}

func (m machineDo) Clauses(conds ...clause.Expression) *machineDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m machineDo) Returning(value interface{}, columns ...string) *machineDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m machineDo) Not(conds ...gen.Condition) *machineDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m machineDo) Or(conds ...gen.Condition) *machineDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m machineDo) Select(conds ...field.Expr) *machineDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m machineDo) Where(conds ...gen.Condition) *machineDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m machineDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *machineDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m machineDo) Order(conds ...field.Expr) *machineDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m machineDo) Distinct(cols ...field.Expr) *machineDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m machineDo) Omit(cols ...field.Expr) *machineDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m machineDo) Join(table schema.Tabler, on ...field.Expr) *machineDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m machineDo) LeftJoin(table schema.Tabler, on ...field.Expr) *machineDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m machineDo) RightJoin(table schema.Tabler, on ...field.Expr) *machineDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m machineDo) Group(cols ...field.Expr) *machineDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m machineDo) Having(conds ...gen.Condition) *machineDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m machineDo) Limit(limit int) *machineDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m machineDo) Offset(offset int) *machineDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m machineDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *machineDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m machineDo) Unscoped() *machineDo {
	return m.withDO(m.DO.Unscoped())
}

func (m machineDo) Create(values ...*model.Machine) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m machineDo) CreateInBatches(values []*model.Machine, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m machineDo) Save(values ...*model.Machine) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m machineDo) First() (*model.Machine, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Machine), nil
	}
}

func (m machineDo) Take() (*model.Machine, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Machine), nil
	}
}

func (m machineDo) Last() (*model.Machine, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Machine), nil
	}
}

func (m machineDo) Find() ([]*model.Machine, error) {
	result, err := m.DO.Find()
	return result.([]*model.Machine), err
}

func (m machineDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Machine, err error) {
	buf := make([]*model.Machine, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m machineDo) FindInBatches(result *[]*model.Machine, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m machineDo) Attrs(attrs ...field.AssignExpr) *machineDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m machineDo) Assign(attrs ...field.AssignExpr) *machineDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m machineDo) Joins(fields ...field.RelationField) *machineDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m machineDo) Preload(fields ...field.RelationField) *machineDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m machineDo) FirstOrInit() (*model.Machine, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Machine), nil
	}
}

func (m machineDo) FirstOrCreate() (*model.Machine, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Machine), nil
	}
}

func (m machineDo) FindByPage(offset int, limit int) (result []*model.Machine, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m machineDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m machineDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m machineDo) Delete(models ...*model.Machine) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *machineDo) withDO(do gen.Dao) *machineDo {
	m.DO = *do.(*gen.DO)
	return m
}
