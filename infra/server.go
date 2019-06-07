package infra

import (
	"net/http"
	"time"

	newrelic "github.com/newrelic/go-agent"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/pressly/chi/middleware"
)

type ServerContext struct {
	Router *chi.Mux
	NRApp  *newrelic.Application
}

type StructuredLogger struct {
	Logger *logrus.Logger
}

func New() *ServerContext {
	router := chi.NewRouter()

	router.Use(middleware.Timeout(3 * time.Second))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Throttle(1000))
	router.Use(middleware.NoCache)
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	defer logrus.Info("server route setup completed")
	return &ServerContext{Router: router}
}

func (ctx *ServerContext) StartServer(address string) {
	http.ListenAndServe(address, ctx.Router)
}

func (ctx *ServerContext) Register(path string, handle func(http.ResponseWriter, *http.Request), method string) {
	ctx.Router.MethodFunc(method, path, handle)
}

func (ctx *ServerContext) Mount(path string, handler http.Handler) {
	ctx.Router.Mount(path, handler)
}
