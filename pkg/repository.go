package pkg

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	GetAll(items *[]T, pagination *PaginationQuery, condition interface{}, relations ...string) error
	GetBy(condition interface{}, item *T, relations ...string) error
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

func (r *baseRepository[T]) GetAll(items *[]T, pagination *PaginationQuery, condition interface{}, relations ...string) error {
	query := r.db.Where(condition)

	if pagination.Page != nil && pagination.PageSize != nil {
		offset := *pagination.Page * *pagination.PageSize
		query = query.Offset(offset).Limit(*pagination.PageSize)
	}

	if pagination.Sort != nil && pagination.Order != nil {
		orderClause := *pagination.Sort + " " + *pagination.Order
		query = query.Order(orderClause)
	}

	for _, relation := range relations {
		query = query.Preload(relation)
	}

	return query.Find(items).Error
}

func (r *baseRepository[T]) GetBy(condition interface{}, item *T, relations ...string) error {
	query := r.db.Where(condition)

	for _, relation := range relations {
		query = query.Preload(relation)
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
