package app

import (
	"example.com/fxdemo/internal/app/httpsrv"
	"example.com/fxdemo/internal/app/service"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			zap.NewProduction,
			NewConfig,
		),

		fx.Options(
			httpsrv.NewModule(),
			service.NewModule(),
		),

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),

		fx.Invoke(func(log *zap.Logger) {
			log.Info("application started")
		}),
	)
}
