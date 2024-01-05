package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func StartDB() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("root:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_PASS, DB_HOST, DB_NAME)
	// dsn := "root:@tcp(0.0.0.0:3306)/portobello?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Erro ao conectar à Database")
	}
	db = database
}

func GetDatabase() *gorm.DB {
	return db
}
