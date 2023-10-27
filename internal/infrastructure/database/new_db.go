package database

import (
	"github.com/danilocordeirodev/go-email/internal/domain/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "host=localhost user=go-email password=go-email dbname=go-email port=15432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("faild to connect to database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})
	return db
}
