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
	ACCRUAL_SYSTEM_ADDRESS string `yaml:"ACCRUAL_SYSTEM_ADDRESS"`
	DATABASE_URI           string `yaml:"DATABASE_URI"`
	RUN_ADDRESS            string `yaml:"RUN_ADDRESS"`
}
type ConfigFromENV struct {
	ACCRUAL_SYSTEM_ADDRESS string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	DATABASE_URI           string `env:"DATABASE_URI"`
	RUN_ADDRESS            string `env:"RUN_ADDRESS"`
}
type ServerConfig interface {
	GetConfig() (config *MainConfig, err error)
}

func GetConfig() (config *MainConfig, err error) {

	var ACCRUAL_SYSTEM_ADDRESS, DATABASE_URI, RUN_ADDRESS, configFileName string
	flag.StringVarP(&configFileName, "config", "c", "./config.yml", "path to the configuration file")
	flag.StringVarP(&ACCRUAL_SYSTEM_ADDRESS, "accrual_system_adders", "r", "", "Accrual system address")
	flag.StringVarP(&DATABASE_URI, "database_uri", "d", "", "Base URL")
	flag.StringVarP(&RUN_ADDRESS, "run_server", "a", "", "Run server")
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	log.Infof("Loading configuration at '%s'", configFileName)
	configFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Errorf("can't open the config file: %s", err)

	}
	// Default values
	config = &MainConfig{
		ACCRUAL_SYSTEM_ADDRESS: "localhost:8080",
		RUN_ADDRESS:            "localhost:8080",
		DATABASE_URI:           "",
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
		if cfgenv.ACCRUAL_SYSTEM_ADDRESS != "" {
			config.ACCRUAL_SYSTEM_ADDRESS = cfgenv.ACCRUAL_SYSTEM_ADDRESS
		}
		if cfgenv.DATABASE_URI != "" {
			config.DATABASE_URI = cfgenv.DATABASE_URI
		}
		if cfgenv.RUN_ADDRESS != "" {
			config.RUN_ADDRESS = cfgenv.RUN_ADDRESS
		}
	}

	if RUN_ADDRESS != "" {
		config.RUN_ADDRESS = RUN_ADDRESS
	}
	if DATABASE_URI != "" {
		config.DATABASE_URI = DATABASE_URI
	}
	if ACCRUAL_SYSTEM_ADDRESS != "" {
		config.ACCRUAL_SYSTEM_ADDRESS = ACCRUAL_SYSTEM_ADDRESS
	}
	//***postgres:5432/praktikum?sslmode=disable
	log.Info(config.DATABASE_URI)
	log.Info("Configuration loaded")
	return
}
