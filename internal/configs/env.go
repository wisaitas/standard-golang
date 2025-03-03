package configs

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Environment struct {
	PORT                   string `env:"PORT" envDefault:"8082"`
	DB_HOST                string `env:"DB_HOST" envDefault:"localhost"`
	DB_USER                string `env:"DB_USER" envDefault:"postgres"`
	DB_PASSWORD            string `env:"DB_PASSWORD" envDefault:"root"`
	DB_NAME                string `env:"DB_NAME" envDefault:"postgres"`
	DB_PORT                string `env:"DB_PORT" envDefault:"8080"`
	JWT_SECRET             string `env:"JWT_SECRET" envDefault:"secret"`
	REDIS_HOST             string `env:"REDIS_HOST" envDefault:"localhost"`
	REDIS_PORT             string `env:"REDIS_PORT" envDefault:"8081"`
	MAX_FILE_SIZE          int64  `env:"MAX_FILE_SIZE" envDefault:"5"`
	DISTRICT_FILE_PATH     string `env:"DISTRICT_FILE_PATH" envDefault:"./data/districts.json"`
	SUB_DISTRICT_FILE_PATH string `env:"SUB_DISTRICT_FILE_PATH" envDefault:"./data/sub_districts.json"`
	PROVINCE_FILE_PATH     string `env:"PROVINCE_FILE_PATH" envDefault:"./data/provinces.json"`
}

var ENV Environment

func LoadEnv() {
	if err := env.Parse(&ENV); err != nil {
		log.Fatalf("error parsing environment variables: %v\n", err)
	}
}
