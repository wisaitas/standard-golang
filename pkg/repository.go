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
	WithTx(tx *gorm.DB) BaseRepository[T]
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: db,
	}
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

func (r *baseRepository[T]) WithTx(tx *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: tx,
	}
}
