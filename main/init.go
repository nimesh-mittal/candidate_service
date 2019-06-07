package main

import (
	"candidate_service/commons"
	"os"
)

func InitEnvVariables() {
	os.Setenv("ENVIRONMENT", commons.PROD)
}
