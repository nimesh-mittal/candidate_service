package config

import (
	"candidate_service/pkg/commons"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewViperConfig() *Configuration {
	viper.SetConfigName(GetConfigName())
	viper.AddConfigPath(commons.ConfigDir)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file, %s", err)
	}

	var configuration Configuration
	err := viper.Unmarshal(&configuration)
	if err != nil {
		logrus.Fatalf("unable to decode into struct, %v", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.Unmarshal(&configuration)
		if err != nil {
			logrus.Fatalf("unable to decode into struct, %v", err)
		}
		logrus.Info("Config file changed:", e.Name)
	})

	return &configuration
}

func GetConfigName() string {
	env := os.Getenv(commons.EnvVariable)
	if env == commons.Dev {
		return commons.DevConfigPath
	} else if env == commons.Qa {
		return commons.QaConfigPath
	} else if env == commons.Prod {
		return commons.ProdConfigPath
	} else {
		logrus.Fatalf("ENVIRONMENT variable is not set")
	}
	return commons.Empty
}

var once sync.Once
var conf *Configuration

func GetInstance() *Configuration {
	once.Do(func() {
		conf = NewViperConfig()
		defer logrus.Info("config setup completed")
	})

	return conf
}
