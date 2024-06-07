package main

import (
	"trainee3/lib/apiCode"
	"trainee3/router"
	"trainee3/wire"

	"github.com/sirupsen/logrus"
)

func main() {
	app, err := wire.InitializeApplication()
	if err != nil {
		logrus.Fatalf("InitializeHandler error: %v", err)
	}
	router.Set(app.Handlers)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	code := app.Services.Refresh.ServicesRefresh()
	if code != apiCode.Success {
		logrus.Fatalf("ServicesRefresh error: %v", code)
	}
}
