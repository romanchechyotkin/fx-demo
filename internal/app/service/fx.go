package service

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"service",
		fx.Provide(),
		fx.Invoke(
			func(lc fx.Lifecycle, log *zap.Logger) {
				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						log.Info("service started")
						return nil
					},
					OnStop: func(_ context.Context) error {
						log.Info("service stopped")
						return nil
					},
				})
			},
		),

		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("service")
		}),
	)
}
