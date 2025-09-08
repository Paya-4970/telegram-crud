package main

import (
	"log"

	"github.com/Paya-4970/telegram-crud/configs"
	"github.com/Paya-4970/telegram-crud/internal/app"
	"github.com/Paya-4970/telegram-crud/internal/database"
	"github.com/Paya-4970/telegram-crud/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := database.InitDB(cfg)

	e := gin.Default()
	a := app.Init(db)
	routers.InitRouter(e, a)

	log.Fatal(e.Run(":" + cfg.AppPort))
}
