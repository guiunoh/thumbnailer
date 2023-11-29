package config

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func Config(config interface{}, name string) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	ext := filepath.Ext(name)
	viper.SetConfigName(strings.TrimSuffix(name, ext))
	viper.SetConfigType(ext[1:])

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
}
