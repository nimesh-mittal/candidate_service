package infra

import (
	_ "candidate_service/docs"

	"github.com/sirupsen/logrus"

	"github.com/pressly/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterSwaggerDocs(serverContext *ServerContext) {

	r := serverContext.Router

	r.With(middleware.SetHeader("Content-Type", "")).Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	defer logrus.Info("swagger setup completed")
}
