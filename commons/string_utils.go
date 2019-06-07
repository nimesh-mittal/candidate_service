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

// New Relic
const (
	NR_APP_NAME = "sample_service1"
	NR_LIC_KEY  = "<new relic key here>"
)

const INVALID_REQUEST_PARAMETER = "INVALID_REQUEST_PARAMETER"
const INVALID_REQUEST_BODY = "INVALID_REQUEST_BODY"

func IsEmpty(str string) bool {
	if len(str) <= 0 {
		return true
	}
	return false
}
