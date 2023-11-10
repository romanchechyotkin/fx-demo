package httpsrv

import (
	"fmt"

	"go.uber.org/config"
)

type HTTPConfig struct {
	Bind        string `yaml:"bind"`
	Port        string `yaml:"port"`
	IdleTimeout string `yaml:"idle_timeout"`
}

func NewConfig(provider config.Provider) (*HTTPConfig, error) {
	var cfg HTTPConfig
	var err error

	if err := provider.Get("http").Populate(&cfg); err != nil {
		err = fmt.Errorf("failed to get http config: %w", err)
	}
	return &cfg, err
}
