package configs

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/wisaitas/standard-golang/internal/env"
	"github.com/wisaitas/standard-golang/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		env.DB_HOST,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_NAME,
		env.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := autoMigrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	if err := autoSeed(db); err != nil {
		log.Fatalf("failed to seed database: %v", err)
	}

	log.Println("database connected successfully")
	return db
}

func autoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Province{},
		&models.District{},
		&models.SubDistrict{},
		&models.Address{},
		&models.UserHistory{},
	); err != nil {
		return fmt.Errorf("error migrating database: %w", err)
	}

	log.Println("database migrated successfully")

	return nil
}

func autoSeed(db *gorm.DB) error {
	seedConfigs := []struct {
		model       interface{}
		filename    string
		destination interface{}
		entityName  string
	}{
		{&models.Province{}, env.PROVINCE_FILE_PATH, &[]models.Province{}, "provinces"},
		{&models.District{}, env.DISTRICT_FILE_PATH, &[]models.District{}, "districts"},
		{&models.SubDistrict{}, env.SUB_DISTRICT_FILE_PATH, &[]models.SubDistrict{}, "sub districts"},
	}

	for _, config := range seedConfigs {
		if err := seedIfEmpty(db, config.model, config.filename, config.destination, config.entityName); err != nil {
			return err
		}
	}

	log.Println("database seeded successfully")
	return nil
}

func seedIfEmpty(db *gorm.DB, model interface{}, filename string, destination interface{}, entityName string) error {
	var count int64
	if err := db.Model(model).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking %s: %w", entityName, err)
	}

	if count == 0 {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("error opening %s file: %v", entityName, err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("error reading %s file: %v", entityName, err)
		}

		if err := json.Unmarshal(byteData, destination); err != nil {
			log.Fatalf("error unmarshaling %s: %v", entityName, err)
		}

		if err := db.CreateInBatches(destination, 100).Error; err != nil {
			return fmt.Errorf("error seeding %s: %w", entityName, err)
		}
	}

	return nil
}
