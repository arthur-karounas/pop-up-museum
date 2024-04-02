package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server      Server
	MySQL       MySQL
	Application Application
	Email       Email
}

type Server struct {
	Port string
	Mode string
}

type MySQL struct {
	Name   string
	Host   string
	Port   string
	DBName string
}

type Application struct {
	Uploads     string
	Appeal      EmailMessage
	Reservation EmailMessage
	Passupdate  EmailMessage
}

type EmailMessage struct {
	Subject string
	Message string
}

type Email struct {
	Host string
	Port int
}

func InitConfig() (*Config, error) {
	viperobj := viper.New()

	viperobj.AddConfigPath("configs")
	viperobj.SetConfigName("config")
	viperobj.SetConfigType("yaml")
	if err := viperobj.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	if err := viperobj.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
