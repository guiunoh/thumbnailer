package database

import (
	"database/sql"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySql(dsn string) Database {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	size := runtime.NumCPU()
	conn.SetMaxIdleConns(size)
	conn.SetMaxOpenConns(size*2 + 2)
	conn.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		Conn:       conn,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		DisableAutomaticPing:   true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return &mysqlDB{db: db}
}

type mysqlDB struct {
	db *gorm.DB
}

// Close implements Database.
func (s *mysqlDB) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

// DB implements Database.
func (s *mysqlDB) DB() *gorm.DB {
	return s.db
}
