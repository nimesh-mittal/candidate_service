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
	viper.AddConfigPath(commons.CONFIG_DIR)

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
	env := os.Getenv(commons.ENV_VARIABLE)
	if env == commons.DEV {
		return commons.DEV_CONFIG_PATH
	} else if env == commons.QA {
		return commons.QA_CONFIG_PATH
	} else if env == commons.PROD {
		return commons.PROD_CONFIG_PATH
	} else {
		logrus.Fatalf("ENVIRONMENT variable is not set")
	}
	return commons.EMPTY
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
