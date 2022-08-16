package main

import (
	"os"

	"github.com/borisbbtest/ya-dr/internal/app"
	config "github.com/borisbbtest/ya-dr/internal/config"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("context", "main")

// готово
func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})

	cfg, err := config.GetConfig()
	if err != nil {
		return
	}
	err = app.New(cfg).Start()
	if err != nil {
		log.Fatal(err)
	}
}
