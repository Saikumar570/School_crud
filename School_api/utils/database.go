package utils

import (
	"fmt"
	"log"

	"School_api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=5152 dbname=users port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.School{}, &models.Classes{}, &models.Class6{}, &models.Class7{}, &models.Class8{}, &models.Student{}, &models.Address{})

	fmt.Println("Database connected and migrated")
	return db
}
