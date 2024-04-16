package config_test

import (
	"testing"

	"github.com/guiunoh/thumbnailer/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// test code
	type _config struct {
		Service struct {
			ID string
		}
	}

	var cfg _config
	config.Config(&cfg, "../../config/config.yaml")
	assert.Equal(t, "thumbnailer", cfg.Service.ID)
}

func TestConfigPanic(t *testing.T) {
	assert.Panics(t, func() {
		type _config struct {
			Profile string
		}

		var cfg _config
		config.Config(&cfg, "config.yaml")
	}, "The code did not panic")
}
