package db

import (
	"fmt"
	"log"
	"time"

	"github.com/Witnot/scraper/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=postgres user=postgres password=postgres dbname=scraper port=5432 sslmode=disable"
	var err error

	// Retry loop (wait for DB)
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Database not ready, retrying... (%d/10)", i+1)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = DB.AutoMigrate(&models.Product{}, &models.PriceRecord{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	fmt.Println("✅ Database connected & migrated")
}
