package main

import (
	"fmt"
	"newApp/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	fmt.Println("sever up & started")
}
