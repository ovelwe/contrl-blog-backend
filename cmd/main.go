package main

import (
	"contrl-blog/configs"
	"contrl-blog/internal/db"
	"contrl-blog/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.LoadCfg()

	db.ConnectDatabase(config)

	db.Migrate(db.DB)

	router := gin.Default()

	routes.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка при загрузке сервака: %v", err)
	}
}
