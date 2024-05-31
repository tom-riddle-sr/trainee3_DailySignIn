package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

type MysqlConfig struct {
	UserName string `env:"Mysql_UserName"`
	Password string `env:"Mysql_Password"`
	Addr     string `env:"Mysql_Addr"`
	Port     int    `env:"Mysql_Port"`
	Database string `env:"Mysql_Database"`
}

type MongoConfig struct {
	ApplyURI string `env:"Mongo_ApplyURI"`
	Database string `env:"Mongo_Database"`
}

type RedisConfig struct {
	Addr     string `env:"Redis_Addr"`
	Password string `env:"Redis_Password"`
	DB       int    `env:"Redis_DB"`
}

type Other struct {
	Mode string `env:"MODE"`
}

type Config struct {
	MysqlConfig MysqlConfig
	MongoConfig MongoConfig
	RedisConfig RedisConfig
	Other       Other
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(&cfg.MysqlConfig); err != nil {
		logrus.Fatal(err)
	}
	if err := env.Parse(&cfg.MongoConfig); err != nil {
		logrus.Fatal(err)
	}
	if err := env.Parse(&cfg.RedisConfig); err != nil {
		logrus.Fatal(err)
	}
	if err := env.Parse(&cfg.Other); err != nil {
		logrus.Fatal(err)
	}
	return cfg, nil
}
