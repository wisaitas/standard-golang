package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabaseSQL(
	host string,
	port string,
	user string,
	password string,
	database string,
	driver string,
) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host,
		user,
		password,
		database,
		port,
	)

	switch driver {
	case "postgres":
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		return db
	case "mysql":
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		return db

	default:
		log.Fatalf("invalid driver: %s", driver)
	}

	return nil
}
