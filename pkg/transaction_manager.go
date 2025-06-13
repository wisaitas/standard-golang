package pkg

import (
	"fmt"

	"gorm.io/gorm"
)

type TransactionManager interface {
	ExecuteInTransaction(fn func(tx *gorm.DB) error) error
	GetTransaction() *gorm.DB
	Begin() error
	Commit() error
	Rollback() error
}

type transactionManager struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewTransactionManager(db *gorm.DB) TransactionManager {
	return &transactionManager{
		db: db,
	}
}

func (tm *transactionManager) ExecuteInTransaction(fn func(tx *gorm.DB) error) error {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("[Share Package TransactionUtil] : %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("[Share Package TransactionUtil] : %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("[Share Package TransactionUtil] : %w", err)
	}

	return nil
}

func (tm *transactionManager) GetTransaction() *gorm.DB {
	if tm.tx != nil {
		return tm.tx
	}
	return tm.db
}

func (tm *transactionManager) Begin() error {
	tx := tm.db.Begin()

	if tx.Error != nil {
		return fmt.Errorf("[Share Package TransactionUtil] : %w", tx.Error)
	}

	tm.tx = tx
	return nil
}

func (tm *transactionManager) Commit() error {
	if tm.tx == nil {
		return nil
	}

	if err := tm.tx.Commit().Error; err != nil {
		return fmt.Errorf("[Share Package TransactionUtil] : %w", err)
	}

	tm.tx = nil
	return nil
}

func (tm *transactionManager) Rollback() error {
	if tm.tx == nil {
		return nil
	}

	if err := tm.tx.Rollback().Error; err != nil {
		return fmt.Errorf("[Share Package TransactionUtil] : %w", err)
	}

	tm.tx = nil
	return nil
}
