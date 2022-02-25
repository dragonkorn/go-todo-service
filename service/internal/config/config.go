package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Configuration struct {
	Port         string `env:"PORT" envDefault:"8000"`
	DbConnection string `env:"DB_CONNECTION" envDefault:"host=localhost user=user password=password dbname=todo port=5432 sslmode=disable"`
	IsDbDebug    bool   `env:"IS_DB_DEBUG" envDefault:"true"`
}

func NewConfiguration() *Configuration {
	cfg := Configuration{}

	if err := env.Parse(&cfg); err != nil {
		fmt.Println(err)
	}

	return &cfg
}
