package repository

import (
	"lhon/postgres-rest/internal/models"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB


func InitDB() {
     err := godotenv.Load(filepath.Join(".", ".env"))
    if err != nil {
		println("errororrrs")
        log.Fatal("Error loading .env file")
    }
    dsn := "host=" + os.Getenv("DB_HOST") +
        " user=" + os.Getenv("DB_USER") +
        " password=" + os.Getenv("DB_PASSWORD") +
        " dbname=" + os.Getenv("DB_NAME") +
        " port=" + os.Getenv("DB_PORT") +
        " sslmode=" + os.Getenv("DB_SSLMODE")


    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    DB = db

	// Auto-migrate the User model
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

    log.Println("Database connected successfully")
}

func CloseDB() {
  // Get the underlying *sql.DB object
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB: ", err)
	}

	// Close the database connection pool
	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close database connection: ", err)
	}

	log.Println("Database connection closed successfully")
}