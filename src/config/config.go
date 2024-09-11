package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Conf struct {
	Db       Db     `mapstructure:"db"`
	LogLevel string `mapstructure:"logLevel"`
	Server   Server `mapstructure:"server"`
}
type Db struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	LogLevel string `mapstructure:"logLevel"`
}
type Server struct {
	Port int
}

var configuration Conf

func Load() Conf {
	logrus.Debug("Loading configuration")

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	logrus.Info("Reading config file")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	conf := Conf{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		logrus.Fatal(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(_ fsnotify.Event) {
		for _, f := range configChangeHandlers {
			f()
		}
	})

	configuration = conf
	logLevel, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
	}
	logrus.SetLevel(logLevel)
	return configuration
}

var (
	configChangeHandlers = []func(){
		func() {
			level := viper.GetString("logLevel")
			if level == "" {
				return
			}
			logLevel, err := logrus.ParseLevel(level)
			if err != nil {

			}
			logrus.SetLevel(logLevel)
		},
	}
)
