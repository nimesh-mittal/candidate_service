package infra

import (
	"net/http"
	"time"

	newrelic "github.com/newrelic/go-agent"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/pressly/chi/middleware"
)

// server context maintain service level state
type ServerContext struct {
	Router *chi.Mux
	NRApp  *newrelic.Application
}

// structured logger holds instance of logrus
type StructuredLogger struct {
	Logger *logrus.Logger
}

// create new server context
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

// Starts server on address
func (ctx *ServerContext) StartServer(address string) {
	http.ListenAndServe(address, ctx.Router)
}

// register handler with a path
func (ctx *ServerContext) Register(path string, handle func(http.ResponseWriter, *http.Request), method string) {
	ctx.Router.MethodFunc(method, path, handle)
}

// mount sub router with a path
func (ctx *ServerContext) Mount(path string, handler http.Handler) {
	ctx.Router.Mount(path, handler)
}
