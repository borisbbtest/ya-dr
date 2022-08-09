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
	Accrual_system_address string `yaml:"ACCRUAL_SYSTEM_ADDRESS"`
	Database_uri           string `yaml:"DATABASE_URI"`
	Run_address            string `yaml:"RUN_ADDRESS"`
}
type ConfigFromENV struct {
	Accrual_system_address string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	Database_uri           string `env:"DATABASE_URI"`
	Run_address            string `env:"RUN_ADDRESS"`
}
type ServerConfig interface {
	GetConfig() (config *MainConfig, err error)
}

func GetConfig() (config *MainConfig, err error) {

	var accrual_system_address, database_uri, run_address, configFileName string
	flag.StringVarP(&configFileName, "config", "c", "./config.yml", "path to the configuration file")
	flag.StringVarP(&accrual_system_address, "accrual_system_adders", "r", "", "Accrual system address")
	flag.StringVarP(&database_uri, "database_uri", "d", "", "Base URL")
	flag.StringVarP(&run_address, "run_server", "a", "", "Run server")
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	log.Infof("Loading configuration at '%s'", configFileName)
	configFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Errorf("can't open the config file: %s", err)

	}
	// Default values
	config = &MainConfig{
		Accrual_system_address: "localhost:8080",
		Run_address:            "localhost:8080",
		Database_uri:           "",
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
		if cfgenv.Accrual_system_address != "" {
			config.Accrual_system_address = cfgenv.Accrual_system_address
		}
		if cfgenv.Database_uri != "" {
			config.Database_uri = cfgenv.Database_uri
		}
		if cfgenv.Run_address != "" {
			config.Run_address = cfgenv.Run_address
		}
	}

	if run_address != "" {
		config.Run_address = run_address
	}
	if database_uri != "" {
		config.Database_uri = database_uri
	}
	if accrual_system_address != "" {
		config.Accrual_system_address = accrual_system_address
	}
	//***postgres:5432/praktikum?sslmode=disable
	log.Info(config.Database_uri)
	log.Info("Configuration loaded")
	return
}
