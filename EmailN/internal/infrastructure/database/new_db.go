package database

import (
	"emailn/internal/domain/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDb() *gorm.DB {
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect to database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
