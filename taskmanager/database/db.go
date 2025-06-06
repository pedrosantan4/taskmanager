package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	Completed   bool
}

func Migrate() {
	DB.AutoMigrate(&Task{})
	fmt.Println("Migration completed!")
}

var DB *gorm.DB

func Connect() {
	dsn := "host=host.docker.internal user=postgres password=postgres dbname=taskdb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Database connected!")
}
