package config

import (
	goflag "flag"

	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

type MainConfig struct {
	AccrualSystemAddress string `yaml:"ACCRUAL_SYSTEM_ADDRESS" env:"ACCRUAL_SYSTEM_ADDRESS"`
	DatabaseURI          string `yaml:"DATABASE_URI" env:"DATABASE_URI"`
	RunAddress           string `yaml:"RUN_ADDRESS" env:"RUN_ADDRESS"`
}
type ServerConfig interface {
	GetConfig() (config *MainConfig, err error)
}

func GetConfig() (config *MainConfig, err error) {
	config = &MainConfig{}

	var log = logrus.WithField("context", "system_loyalty")
	flag.StringVarP(&config.AccrualSystemAddress, "accrual_system_adders", "r", config.AccrualSystemAddress, "Accrual system address")
	flag.StringVarP(&config.DatabaseURI, "database_uri", "d", config.DatabaseURI, "Base URL")
	flag.StringVarP(&config.RunAddress, "run_server", "a", config.RunAddress, "Run server")
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	err = env.Parse(config)
	if err != nil {
		log.Errorf("can't start the listening thread: %s", err)
		return
	}

	//***postgres:5432/praktikum?sslmode=disable
	log.Info("Configuration loaded")
	return
}
