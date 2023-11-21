package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSqlite(dsn string) Database {
	var err error
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("db connect failed")
	}

	return &sqliteDB{db: db, dsn: dsn}
}

type sqliteDB struct {
	db  *gorm.DB
	dsn string
}

// Close implements Database.
func (s *sqliteDB) Close() error {
	if sqlDB, err := s.db.DB(); err == nil {
		sqlDB.Close()
	}

	return os.Remove(s.dsn)
}

func (s *sqliteDB) DB() *gorm.DB {
	return s.db
}
