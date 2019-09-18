package commons

const (
	EMPTY        = ""
	ENV_VARIABLE = "ENVIRONMENT"
	OK           = "ok"
	TrackingID   = "TrackingID"
)

// Config Viper
const (
	DEV              = "DEV"
	QA               = "QA"
	PROD             = "PROD"
	DEV_CONFIG_PATH  = "config_dev"
	QA_CONFIG_PATH   = "config_qa"
	PROD_CONFIG_PATH = "config"
	CONFIG_DIR       = "./config"
)

// Heartbeat Status
const (
	GREEN_STATUS = "green"
	AMBER_STATUS = "amber"
	RED_STATUS   = "red"
)

// Mongo
const (
	CANDIDATE_COLL = "candidates"
	CANDIDATE_DB   = "core_db"
)

const INVALID_REQUEST_PARAMETER = "INVALID_REQUEST_PARAMETER"
const INVALID_REQUEST_BODY = "INVALID_REQUEST_BODY"
