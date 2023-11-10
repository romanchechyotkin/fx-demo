package httpsrv

import (
	"example.com/fxdemo/internal/app/httpsrv/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	srv    *http.Server
	log    *zap.Logger
	router *chi.Mux
}

func NewServer(log *zap.Logger) *Server {
	router := chi.NewRouter()

	srv := &Server{
		srv:    &http.Server{Addr: ":8080", Handler: router},
		log:    log,
		router: router,
	}

	srv.router.Use(middleware.AllowContentType("application/json"))

	//srv.router.Group(func(statusGrp chi.Router) {
	//	statusGrp.Handle("/metrics", index())
	//})

	srv.router.Get("/health", handlers.Health())

	return srv
}
