package database

import (
	"emailn/internal/domain/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=emailn_dev password=123456 dbname=emailn_dev port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect to database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
