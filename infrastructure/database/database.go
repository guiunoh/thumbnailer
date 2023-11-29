package database

import (
	"fmt"

	"gorm.io/gorm"
)

type Config struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Charset  string
	TimeZone string
}

func NewDatabase(config Config) Database {
	switch config.Driver {
	case "sqlite3":
		dsn := fmt.Sprintf("%s.db", config.Name)
		return NewSqlite(dsn)
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.Name,
			config.Charset,
			true,
			config.TimeZone,
		)
		return NewMySql(dsn)
	default:
		panic("unknown database driver: " + config.Driver)
	}
}

type Database interface {
	DB() *gorm.DB
	Close() error
}
