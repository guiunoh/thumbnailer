package main

import (
	"flag"
	"thumbnailer/infrastructure/database"
	"thumbnailer/pkg/config"

	"github.com/spf13/viper"
)

type Config struct {
	Profile string
	Service struct {
		Port int
	}
	Monitor struct {
		Port int
		Path string
	}
	Log struct {
		Level      string
		Path       string
		MaxSize    int
		MaxBackups int
		MaxAge     int
	}
	DB database.Config
}

func loadConfig() Config {
	var name string
	flag.StringVar(&name, "config", "./config/config.yaml", "config file name. --config=config.yaml")
	flag.Parse()

	viper.SetDefault("profile", "dev")
	viper.SetDefault("service.port", 8080)
	viper.SetDefault("monitor.port", 9090)
	viper.SetDefault("monitor.path", "metrics")
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.path", "./logs/app.log")
	viper.SetDefault("log.maxSize", 100)
	viper.SetDefault("log.maxBackups", 10)
	viper.SetDefault("log.maxAge", 30)
	viper.SetDefault("db.driver", "sqlite3")
	viper.SetDefault("db.name", "thumbnailer")
	viper.SetDefault("db.charset", "utf8mb4")
	viper.SetDefault("db.timeZone", "Local")

	var cfg Config
	config.Config(&cfg, name)
	return cfg
}
