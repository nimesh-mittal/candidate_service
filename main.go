package main

import (
	"candidate_service/infra"
	"candidate_service/pkg/admin"
	"candidate_service/pkg/candidate"
	"candidate_service/pkg/commons"
	"os"
)

func Init() {
	// TODO: move this to docker file
	os.Setenv("ENVIRONMENT", commons.PROD)
}

func main() {

	// initialise application
	Init()
	serverContext := infra.New()

	// Register Infra Components
	infra.RegisterSwaggerDocs(serverContext)
	infra.RegisterNewrelic()

	// Register Routes
	candidateResourceContext := candidate.NewCandidateResource()
	defer candidateResourceContext.SafeClose()

	adminResourceContext := admin.NewAdminResource()
	defer adminResourceContext.SafeClose()

	serverContext.Mount("/api/v1/candidates", candidateResourceContext.NewCandidateRouter())
	serverContext.Mount("/api/v1/admin", adminResourceContext.NewAdminRouter())

	// Start server
	serverContext.StartServer(":3000")
}
