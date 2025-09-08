package database

import (
	"log"

	"github.com/Paya-4970/telegram-crud/configs"
	"github.com/Paya-4970/telegram-crud/internal/models"
	_ "github.com/Paya-4970/telegram-crud/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg config.Config) *gorm.DB {
	var err error
	DB, err = gorm.Open(mysql.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	// AutoMigrate مدل‌ها
	if err := DB.AutoMigrate(
		&models.Food{},
	); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}
	return DB
}
