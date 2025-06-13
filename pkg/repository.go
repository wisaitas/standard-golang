package pkg

import (
	"gorm.io/gorm"
)

type Condition struct {
	Query interface{}
	Args  []interface{}
}

func NewCondition(query interface{}, args ...interface{}) *Condition {
	return &Condition{
		Query: query,
		Args:  args,
	}
}

type Relation struct {
	Query string
	Args  []interface{}
}

func NewRelation(query string, args ...interface{}) *Relation {
	return &Relation{
		Query: query,
		Args:  args,
	}
}

type TxManager struct {
	tx *gorm.DB
}

func NewTxManager(db *gorm.DB) *TxManager {
	return &TxManager{
		tx: db.Begin(),
	}
}

func (tm *TxManager) GetTx() *gorm.DB {
	return tm.tx
}

func (tm *TxManager) Commit() error {
	return tm.tx.Commit().Error
}

func (tm *TxManager) Rollback() error {
	return tm.tx.Rollback().Error
}

type BaseRepository[T any] interface {
	GetAll(items *[]T, pagination *PaginationQuery, condition *Condition, relations *[]Relation) error
	GetBy(item *T, condition *Condition, relations *[]Relation) error
	Create(item *T) error
	CreateMany(items *[]T) error
	Update(item *T) error
	UpdateMany(items *[]T) error
	Save(item *T) error
	SaveMany(items *[]T) error
	Delete(item *T) error
	DeleteMany(items *[]T) error
	WithTxManager(tm *TxManager) BaseRepository[T]
	GetDB() *gorm.DB
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: db,
	}
}

func (r *baseRepository[T]) GetDB() *gorm.DB {
	return r.db
}

func (r *baseRepository[T]) GetAll(items *[]T, pagination *PaginationQuery, condition *Condition, relations *[]Relation) error {
	query := r.db

	if condition != nil {
		query = query.Where(condition.Query, condition.Args...)
	}

	if pagination != nil {
		if pagination.Page != nil && pagination.PageSize != nil {
			offset := *pagination.Page * *pagination.PageSize
			query = query.Offset(offset).Limit(*pagination.PageSize)
		}

		if pagination.Sort != nil && pagination.Order != nil {
			orderClause := *pagination.Sort + " " + *pagination.Order
			query = query.Order(orderClause)
		}
	}

	if relations != nil {
		for _, relation := range *relations {
			query = query.Preload(relation.Query, relation.Args...)
		}
	}

	return query.Find(items).Error
}

func (r *baseRepository[T]) GetBy(item *T, condition *Condition, relations *[]Relation) error {
	query := r.db

	if condition != nil {
		query = query.Where(condition.Query, condition.Args...)
	}

	if relations != nil {
		for _, relation := range *relations {
			query = query.Preload(relation.Query, relation.Args...)
		}
	}

	return query.First(item).Error
}

func (r *baseRepository[T]) Create(item *T) error {
	return r.db.Create(item).Error
}

func (r *baseRepository[T]) CreateMany(items *[]T) error {
	return r.db.Create(items).Error
}

func (r *baseRepository[T]) Update(item *T) error {
	return r.db.Updates(item).Error
}

func (r *baseRepository[T]) UpdateMany(items *[]T) error {
	return r.db.Updates(items).Error
}

func (r *baseRepository[T]) Save(item *T) error {
	return r.db.Save(item).Error
}

func (r *baseRepository[T]) SaveMany(items *[]T) error {
	return r.db.Save(items).Error
}

func (r *baseRepository[T]) Delete(item *T) error {
	return r.db.Delete(item).Error
}

func (r *baseRepository[T]) DeleteMany(items *[]T) error {
	return r.db.Delete(items).Error
}

func (r *baseRepository[T]) WithTxManager(tm *TxManager) BaseRepository[T] {
	return &baseRepository[T]{
		db: tm.GetTx(),
	}
}

func (r *baseRepository[T]) Begin() BaseRepository[T] {
	return &baseRepository[T]{
		db: r.db.Begin(),
	}
}

func (r *baseRepository[T]) Rollback() error {
	return r.db.Rollback().Error
}

func (r *baseRepository[T]) Commit() error {
	return r.db.Commit().Error
}
