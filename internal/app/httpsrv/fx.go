package httpsrv

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"http server",
		fx.Provide(
			NewConfig,
			NewServer,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, srv *Server, log *zap.Logger) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							log.Info("http-server listen and serve", zap.String("on", srv.srv.Addr))
							if err := srv.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
								log.Error("failed to listen and serve", zap.Error(err))
							}
						}()

						return nil
					},
					OnStop: func(ctx context.Context) error {
						return srv.srv.Shutdown(ctx)
					},
				})
			},
		),

		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("http server")
		}),
	)
}
