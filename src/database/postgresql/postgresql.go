package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// TODO read from environment variables
	dsn := "host=localhost user=rockbooks password=rockbooks dbname=rockbooks port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("auto migration has been ran successfully ✅")
	DB = db
	fmt.Println("datbase has been initialized ✅")
	return db
}
