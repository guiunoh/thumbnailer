package main

import (
	"flag"
	_database "thumbnailer/infrastructure/database"
	_config "thumbnailer/pkg/config"
)

type config struct {
	Profile string
	Service struct {
		Port int
	}
	DB _database.Config
}

func Config() config {
	var name string
	flag.StringVar(&name, "config", "./config.yaml", "config file name. --config=config.yaml")
	flag.Parse()

	var cfg config
	_config.Config(&cfg, name)

	return cfg
}
