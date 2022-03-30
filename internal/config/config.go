package config

import (
	"html/template"
	"os"
	"sync"

	"github.com/alexedwards/scs/v2"

	"github.com/Elren44/elog"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Listen struct {
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	}
	Storage        StorageConfig `yaml:"storage"`
	TemplatesCache map[string]*template.Template
	UseCache       bool
	Session        *scs.SessionManager
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func GetConfig() *Config {
	var conf *Config
	var once sync.Once
	cfgPath, ok := os.LookupEnv("CONF_PATH")
	if ok {
		once.Do(func() {
			logger := elog.InitLogger(elog.JsonOutput)

			logger.Info("read configuration application")
			conf = &Config{}
			if err := cleanenv.ReadConfig(cfgPath, conf); err != nil {
				help, _ := cleanenv.GetDescription(conf, nil)
				logger.Info(help)
				logger.Fatal(err)
			}

		})
	}

	return conf
}
