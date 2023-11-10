package app

import (
	"fmt"

	"go.uber.org/config"
	"go.uber.org/fx"
)

type ResultConfig struct {
	fx.Out

	Provider config.Provider
}

func NewConfig() (ResultConfig, error) {
	loader, err := config.NewYAML(config.File("config.yaml"))
	if err != nil {
		return ResultConfig{}, fmt.Errorf("failed to load config: %w", err)
	}

	return ResultConfig{
		Provider: loader,
	}, nil
}
