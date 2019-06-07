package config

type Configuration struct {
	Database Database
	Server   Server
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
