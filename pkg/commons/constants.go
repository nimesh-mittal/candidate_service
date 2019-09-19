package commons

const (
	Empty       = ""
	EnvVariable = "ENVIRONMENT"
	Ok          = "ok"
	TrackingID  = "TrackingID"
)

// Config Viper
const (
	Dev            = "Dev"
	Qa             = "Qa"
	Prod           = "Prod"
	DevConfigPath  = "config_dev"
	QaConfigPath   = "config_qa"
	ProdConfigPath = "config"
	ConfigDir      = "./config"
)

// Heartbeat Status
const (
	GreenStatus = "green"
	AmberStatus = "amber"
	RedStatus   = "red"
)

// Mongo
const (
	CandidateColl = "candidates"
	CandidateDb   = "core_db"
)

const InvalidRequestParameter = "InvalidRequestParameter"
const InvalidRequestBody = "InvalidRequestBody"
