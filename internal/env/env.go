package env

import (
	"log"

	"github.com/caarlos0/env/v11"
)

var (
	PORT                   string
	DB_HOST                string
	DB_USER                string
	DB_PASSWORD            string
	DB_NAME                string
	DB_PORT                string
	JWT_SECRET             string
	REDIS_HOST             string
	REDIS_PORT             string
	MAX_FILE_SIZE          int64
	DISTRICT_FILE_PATH     string
	SUB_DISTRICT_FILE_PATH string
	PROVINCE_FILE_PATH     string
)

type environment struct {
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

var dependency environment

func LoadEnv() {
	if err := env.Parse(&dependency); err != nil {
		log.Fatalf("error parsing environment variables: %v\n", err)
	}

	PORT = dependency.PORT
	DB_HOST = dependency.DB_HOST
	DB_USER = dependency.DB_USER
	DB_PASSWORD = dependency.DB_PASSWORD
	DB_NAME = dependency.DB_NAME
	DB_PORT = dependency.DB_PORT
	JWT_SECRET = dependency.JWT_SECRET
	REDIS_HOST = dependency.REDIS_HOST
	REDIS_PORT = dependency.REDIS_PORT
	MAX_FILE_SIZE = dependency.MAX_FILE_SIZE
	DISTRICT_FILE_PATH = dependency.DISTRICT_FILE_PATH
	SUB_DISTRICT_FILE_PATH = dependency.SUB_DISTRICT_FILE_PATH
	PROVINCE_FILE_PATH = dependency.PROVINCE_FILE_PATH
}
