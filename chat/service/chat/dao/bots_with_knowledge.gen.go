// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"chat/service/chat/model"
)

func newBotsWithKnowledge(db *gorm.DB, opts ...gen.DOOption) botsWithKnowledge {
	_botsWithKnowledge := botsWithKnowledge{}

	_botsWithKnowledge.botsWithKnowledgeDo.UseDB(db, opts...)
	_botsWithKnowledge.botsWithKnowledgeDo.UseModel(&model.BotsWithKnowledge{})

	tableName := _botsWithKnowledge.botsWithKnowledgeDo.TableName()
	_botsWithKnowledge.ALL = field.NewAsterisk(tableName)
	_botsWithKnowledge.ID = field.NewInt64(tableName, "id")
	_botsWithKnowledge.BotID = field.NewInt64(tableName, "bot_id")
	_botsWithKnowledge.KnowledgeID = field.NewInt64(tableName, "knowledge_id")
	_botsWithKnowledge.CreatedAt = field.NewTime(tableName, "created_at")
	_botsWithKnowledge.UpdatedAt = field.NewTime(tableName, "updated_at")

	_botsWithKnowledge.fillFieldMap()

	return _botsWithKnowledge
}

type botsWithKnowledge struct {
	botsWithKnowledgeDo botsWithKnowledgeDo

	ALL         field.Asterisk
	ID          field.Int64 // 机器人初始设置ID
	BotID       field.Int64 // 机器人ID 关联 bots.id
	KnowledgeID field.Int64 // 知识库ID 关联 knowledge.id
	CreatedAt   field.Time  // 创建时间
	UpdatedAt   field.Time  // 更新时间

	fieldMap map[string]field.Expr
}

func (b botsWithKnowledge) Table(newTableName string) *botsWithKnowledge {
	b.botsWithKnowledgeDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b botsWithKnowledge) As(alias string) *botsWithKnowledge {
	b.botsWithKnowledgeDo.DO = *(b.botsWithKnowledgeDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *botsWithKnowledge) updateTableName(table string) *botsWithKnowledge {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.BotID = field.NewInt64(table, "bot_id")
	b.KnowledgeID = field.NewInt64(table, "knowledge_id")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")

	b.fillFieldMap()

	return b
}

func (b *botsWithKnowledge) WithContext(ctx context.Context) *botsWithKnowledgeDo {
	return b.botsWithKnowledgeDo.WithContext(ctx)
}

func (b botsWithKnowledge) TableName() string { return b.botsWithKnowledgeDo.TableName() }

func (b botsWithKnowledge) Alias() string { return b.botsWithKnowledgeDo.Alias() }

func (b botsWithKnowledge) Columns(cols ...field.Expr) gen.Columns {
	return b.botsWithKnowledgeDo.Columns(cols...)
}

func (b *botsWithKnowledge) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *botsWithKnowledge) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 5)
	b.fieldMap["id"] = b.ID
	b.fieldMap["bot_id"] = b.BotID
	b.fieldMap["knowledge_id"] = b.KnowledgeID
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
}

func (b botsWithKnowledge) clone(db *gorm.DB) botsWithKnowledge {
	b.botsWithKnowledgeDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b botsWithKnowledge) replaceDB(db *gorm.DB) botsWithKnowledge {
	b.botsWithKnowledgeDo.ReplaceDB(db)
	return b
}

type botsWithKnowledgeDo struct{ gen.DO }

func (b botsWithKnowledgeDo) Debug() *botsWithKnowledgeDo {
	return b.withDO(b.DO.Debug())
}

func (b botsWithKnowledgeDo) WithContext(ctx context.Context) *botsWithKnowledgeDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b botsWithKnowledgeDo) ReadDB() *botsWithKnowledgeDo {
	return b.Clauses(dbresolver.Read)
}

func (b botsWithKnowledgeDo) WriteDB() *botsWithKnowledgeDo {
	return b.Clauses(dbresolver.Write)
}

func (b botsWithKnowledgeDo) Session(config *gorm.Session) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Session(config))
}

func (b botsWithKnowledgeDo) Clauses(conds ...clause.Expression) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b botsWithKnowledgeDo) Returning(value interface{}, columns ...string) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b botsWithKnowledgeDo) Not(conds ...gen.Condition) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b botsWithKnowledgeDo) Or(conds ...gen.Condition) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b botsWithKnowledgeDo) Select(conds ...field.Expr) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b botsWithKnowledgeDo) Where(conds ...gen.Condition) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b botsWithKnowledgeDo) Order(conds ...field.Expr) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b botsWithKnowledgeDo) Distinct(cols ...field.Expr) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b botsWithKnowledgeDo) Omit(cols ...field.Expr) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b botsWithKnowledgeDo) Join(table schema.Tabler, on ...field.Expr) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b botsWithKnowledgeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *botsWithKnowledgeDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b botsWithKnowledgeDo) RightJoin(table schema.Tabler, on ...field.Expr) *botsWithKnowledgeDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b botsWithKnowledgeDo) Group(cols ...field.Expr) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b botsWithKnowledgeDo) Having(conds ...gen.Condition) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b botsWithKnowledgeDo) Limit(limit int) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b botsWithKnowledgeDo) Offset(offset int) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b botsWithKnowledgeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b botsWithKnowledgeDo) Unscoped() *botsWithKnowledgeDo {
	return b.withDO(b.DO.Unscoped())
}

func (b botsWithKnowledgeDo) Create(values ...*model.BotsWithKnowledge) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b botsWithKnowledgeDo) CreateInBatches(values []*model.BotsWithKnowledge, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b botsWithKnowledgeDo) Save(values ...*model.BotsWithKnowledge) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b botsWithKnowledgeDo) First() (*model.BotsWithKnowledge, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BotsWithKnowledge), nil
	}
}

func (b botsWithKnowledgeDo) Take() (*model.BotsWithKnowledge, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BotsWithKnowledge), nil
	}
}

func (b botsWithKnowledgeDo) Last() (*model.BotsWithKnowledge, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BotsWithKnowledge), nil
	}
}

func (b botsWithKnowledgeDo) Find() ([]*model.BotsWithKnowledge, error) {
	result, err := b.DO.Find()
	return result.([]*model.BotsWithKnowledge), err
}

func (b botsWithKnowledgeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BotsWithKnowledge, err error) {
	buf := make([]*model.BotsWithKnowledge, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b botsWithKnowledgeDo) FindInBatches(result *[]*model.BotsWithKnowledge, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b botsWithKnowledgeDo) Attrs(attrs ...field.AssignExpr) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b botsWithKnowledgeDo) Assign(attrs ...field.AssignExpr) *botsWithKnowledgeDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b botsWithKnowledgeDo) Joins(fields ...field.RelationField) *botsWithKnowledgeDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b botsWithKnowledgeDo) Preload(fields ...field.RelationField) *botsWithKnowledgeDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b botsWithKnowledgeDo) FirstOrInit() (*model.BotsWithKnowledge, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BotsWithKnowledge), nil
	}
}

func (b botsWithKnowledgeDo) FirstOrCreate() (*model.BotsWithKnowledge, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BotsWithKnowledge), nil
	}
}

func (b botsWithKnowledgeDo) FindByPage(offset int, limit int) (result []*model.BotsWithKnowledge, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b botsWithKnowledgeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b botsWithKnowledgeDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b botsWithKnowledgeDo) Delete(models ...*model.BotsWithKnowledge) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *botsWithKnowledgeDo) withDO(do gen.Dao) *botsWithKnowledgeDo {
	b.DO = *do.(*gen.DO)
	return b
}
