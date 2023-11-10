package service

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"service",
		fx.Provide(
			NewService,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, log *zap.Logger, srv *Service) {
				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						log.Info("service started")
						go srv.Run()
						return nil
					},
					OnStop: func(_ context.Context) error {
						log.Info("service stopped")

						srv.worker <- struct{}{}

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
