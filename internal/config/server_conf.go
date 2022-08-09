package config

import (
	goflag "flag"
	"io/ioutil"

	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

var log = logrus.WithField("context", "system_loyalty")

type MainConfig struct {
	AccrualSystemAddress string `yaml:"ACCRUAL_SYSTEM_ADDRESS"`
	DatabaseURI          string `yaml:"DATABASE_URI"`
	RunAddress           string `yaml:"RUN_ADDRESS"`
}
type ConfigFromENV struct {
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	DatabaseURI          string `env:"DATABASE_URI"`
	RunAddress           string `env:"RUN_ADDRESS"`
}
type ServerConfig interface {
	GetConfig() (config *MainConfig, err error)
}

func GetConfig() (config *MainConfig, err error) {

	var accrualSystemAddress, databaseURI, runAddress, configFileName string
	flag.StringVarP(&configFileName, "config", "c", "./config.yml", "path to the configuration file")
	flag.StringVarP(&accrualSystemAddress, "accrual_system_adders", "r", "", "Accrual system address")
	flag.StringVarP(&databaseURI, "database_uri", "d", "", "Base URL")
	flag.StringVarP(&runAddress, "run_server", "a", "", "Run server")
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	log.Infof("Loading configuration at '%s'", configFileName)
	configFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Errorf("can't open the config file: %s", err)

	}
	// Default values
	config = &MainConfig{
		AccrualSystemAddress: "localhost:8080",
		RunAddress:           "localhost:8080",
		DatabaseURI:          "",
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Errorf("can't read the config file: %s", err)
	}

	var cfgenv ConfigFromENV
	err = env.Parse(&cfgenv)
	if err != nil {
		log.Errorf("can't start the listening thread: %s", err)
	} else {
		if cfgenv.AccrualSystemAddress != "" {
			config.AccrualSystemAddress = cfgenv.AccrualSystemAddress
		}
		if cfgenv.DatabaseURI != "" {
			config.DatabaseURI = cfgenv.DatabaseURI
		}
		if cfgenv.RunAddress != "" {
			config.RunAddress = cfgenv.RunAddress
		}
	}

	if runAddress != "" {
		config.RunAddress = runAddress
	}
	if databaseURI != "" {
		config.DatabaseURI = databaseURI
	}
	if accrualSystemAddress != "" {
		config.AccrualSystemAddress = accrualSystemAddress
	}
	//***postgres:5432/praktikum?sslmode=disable
	log.Info("Configuration loaded")
	return
}
