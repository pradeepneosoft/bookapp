package config

import (
	"fmt"
	"log"
	"newApp/models"
	"os"

	//"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env")
	}
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}
	log.Println("database connected ")
	db.AutoMigrate(&models.User{}, &models.Book{})
	return db
}
func ClosedatabaseConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("failed to close")
	}
	dbSql.Close()
}
