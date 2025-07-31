package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Environment variable DB_UR_ is not set")
	}
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed zto connect to DB: %s", err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get raw DB: %s", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %s", err)
	}
	log.Println("Connected to DB successfully")
}
