package db

import (
	"contrl-blog/internal/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("ошибка миграции бд: %v", err)
	}
}
