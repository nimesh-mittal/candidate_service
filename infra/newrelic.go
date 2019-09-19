package infra

import (
	config2 "candidate_service/config"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"

	newrelic "github.com/newrelic/go-agent"
)

// register newrelic with the service name
func RegisterNewrelic() *newrelic.Application {
	config := newrelic.NewConfig(config2.GetInstance().NewRelic.AppName, config2.GetInstance().NewRelic.LicKey)
	app, err := newrelic.NewApplication(config)

	if err != nil {
		logrus.Fatal(err)
	}

	return &app
}

// wrap function for new relic metric reporting
func WrapNR(pattern string, handler func(http.ResponseWriter, *http.Request)) (string, func(http.ResponseWriter, *http.Request)) {
	return newrelic.WrapHandleFunc(*GetNRApp(), pattern, handler)
}

var once sync.Once
var app *newrelic.Application

// init new relic instance once
func GetNRApp() *newrelic.Application {
	once.Do(func() {
		app = RegisterNewrelic()
		defer logrus.Info("newrelic app setup completed")
	})

	return app
}

// start new relic transaction
func StartTx(name string) *newrelic.Transaction {
	txn := (*GetNRApp()).StartTransaction(name, nil, nil)
	return &txn
}

// end new relic transaction
func EndTx(tx *newrelic.Transaction) {
	(*tx).End()
}
