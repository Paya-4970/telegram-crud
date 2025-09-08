package app

import (
	"github.com/Paya-4970/telegram-crud/internal/handlers"
	"github.com/Paya-4970/telegram-crud/internal/repository"
	"github.com/Paya-4970/telegram-crud/internal/services"
	"gorm.io/gorm"
)

type App struct {
	FoodHnadlers *handlers.FoodHandlers
}

func Init(db *gorm.DB) *App{
	// FOOD INIT
	frepo := repository.NewFoodRepo(db)
	fServ := services.NewFoodServices(frepo)
	fHand := handlers.NewFoodServies(fServ)

	return &App{
		FoodHnadlers: fHand,
	}
}