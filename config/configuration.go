package config

type Configuration struct {
	Database Database
	Server   Server
	NewRelic NewRelic
}

type Database struct {
	URL      string
	Name     string
	User     string
	Password string
	Dialect  string
	MongoURL string
}

type Server struct {
	Host string
	Port int
}

type NewRelic struct {
	AppName string
	LicKey  string
}
