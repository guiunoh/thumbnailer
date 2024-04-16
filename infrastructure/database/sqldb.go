package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQL(dsn string, batchSize int) *gorm.DB {
	sqldb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		CreateBatchSize:        batchSize,
	})
	if err != nil {
		panic("db connect failed")
	}

	return sqldb
}
