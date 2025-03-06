package db

import (
	"go-echo/config"
	"go-echo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database with postgresql
func Init() {
	config := config.GetConfig()

	dsn := "host=" + config.DB_HOST + " user=" + config.DB_USER + " password=" + config.DB_PASS + " dbname=" + config.DB_NAME + " port=" + config.DB_PORT + " sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}

	// Migrate the schema (this will auto-create the tables)
	DB.AutoMigrate(&models.User{})
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return DB
}
