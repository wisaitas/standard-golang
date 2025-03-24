package pkg

import (
	"gorm.io/gorm"
)

type TransactionUtil interface {
	ExecuteInTransaction(fn func(tx *gorm.DB) error) error
	GetTransaction() *gorm.DB
	Begin() error
	Commit() error
	Rollback() error
}

type transactionUtil struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewTransactionUtil(db *gorm.DB) TransactionUtil {
	return &transactionUtil{
		db: db,
	}
}

func (tm *transactionUtil) ExecuteInTransaction(fn func(tx *gorm.DB) error) error {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (tm *transactionUtil) GetTransaction() *gorm.DB {
	if tm.tx != nil {
		return tm.tx
	}
	return tm.db
}

func (tm *transactionUtil) Begin() error {
	tx := tm.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	tm.tx = tx
	return nil
}

func (tm *transactionUtil) Commit() error {
	if tm.tx == nil {
		return nil
	}
	err := tm.tx.Commit().Error
	tm.tx = nil
	return err
}

func (tm *transactionUtil) Rollback() error {
	if tm.tx == nil {
		return nil
	}
	err := tm.tx.Rollback().Error
	tm.tx = nil
	return err
}
