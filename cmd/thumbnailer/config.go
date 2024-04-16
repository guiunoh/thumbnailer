package main

import (
	"flag"

	_config "github.com/guiunoh/thumbnailer/pkg/config"
	"github.com/spf13/viper"
)

type Config struct {
	Service struct {
		ID string
	}

	Server struct {
		Port int
		Ping string
	}

	Monitor struct {
		Port int
		Path string
	}

	SqlDB struct {
		DSN       string
		BatchSize int
	}

	Log struct {
		Level      string
		Path       string
		MaxSize    int
		MaxBackups int
		MaxAge     int
	}
}

func config() Config {
	var name string
	flag.StringVar(&name, "config", "./config/config.yaml", "config file name. --config=config.yaml")
	flag.Parse()

	viper.SetDefault("service.id", "thumbnailer")

	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.ping", "/ping")

	viper.SetDefault("monitor.port", 9090)
	viper.SetDefault("monitor.path", "metrics")

	viper.SetDefault("sqldb.dsn", "sync.db")
	viper.SetDefault("sqldb.batchSize", 100)

	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.path", "./logs/app.log")
	viper.SetDefault("log.maxSize", 100)
	viper.SetDefault("log.maxBackups", 10)
	viper.SetDefault("log.maxAge", 30)

	var cfg Config
	_config.Config(&cfg, name)
	return cfg
}
