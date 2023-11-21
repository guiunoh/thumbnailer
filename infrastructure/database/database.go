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
	TimeZone string
}

func NewDatabase(cfg Config) Database {
	switch cfg.Driver {
	case "sqlite3":
		dsn := fmt.Sprintf("%s.db", cfg.Name)
		return NewSqlite(dsn)
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%t&loc=%s",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Name,
			true,
			cfg.TimeZone,
		)
		return NewMySql(dsn)
	default:
		panic("unknown database driver: " + cfg.Driver)
	}
}

type Database interface {
	DB() *gorm.DB
	Close() error
}
