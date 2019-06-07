package main

import (
	"candidate_service/commons"
	"candidate_service/handlers"
	"candidate_service/infra"
	"os"
)

func main() {

	// initialise application
	os.Setenv("ENVIRONMENT", commons.PROD)
	serverContext := infra.New()

	// Register Infra Components
	infra.RegisterSwaggerDocs(serverContext)
	infra.RegisterNewrelic()

	// Register Routes
	candidateResourceContext := handlers.NewCandidateResource()
	defer candidateResourceContext.SafeClose()

	adminResourceContext := handlers.NewAdminResource()
	defer adminResourceContext.SafeClose()

	serverContext.Mount("/api/v1/candidates", candidateResourceContext.NewCandidateRouter())
	serverContext.Mount("/api/v1/admin", adminResourceContext.NewAdminRouter())

	// Start server
	serverContext.StartServer(":3000")
}
