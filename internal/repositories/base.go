package repositories

import (
	"github.com/wisaitas/standard-golang/internal/dtos/request"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	GetAll(items *[]T, pagination *request.PaginationParam) error
	GetById(id uuid.UUID, item *T) error
	Create(item *T) error
	Updates(item *T) error
	Save(item *T) error
	Delete(item *T) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: db,
	}
}

func (r *baseRepository[T]) GetAll(items *[]T, pagination *request.PaginationParam) error {
	if pagination.Page != nil && pagination.PageSize != nil {
		return r.db.Offset(*pagination.Page).Limit(*pagination.PageSize).Find(items).Error
	}
	return r.db.Find(items).Error
}

func (r *baseRepository[T]) GetById(id uuid.UUID, item *T) error {
	return r.db.First(item, id).Error
}

func (r *baseRepository[T]) Create(item *T) error {
	return r.db.Create(item).Error
}

func (r *baseRepository[T]) Updates(item *T) error {
	return r.db.Updates(item).Error
}

func (r *baseRepository[T]) Save(item *T) error {
	return r.db.Save(item).Error
}

func (r *baseRepository[T]) Delete(item *T) error {
	return r.db.Delete(item).Error
}
