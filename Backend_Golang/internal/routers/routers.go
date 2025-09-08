package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Paya-4970/telegram-crud/internal/app"
)

func InitRouter(e *gin.Engine, a *app.App) {
	foodGroup := e.Group("/food") // Grouping food-related routes
	{
		foodGroup.GET("/get/:id", a.FoodHnadlers.Get)
		foodGroup.GET("/get/:id", a.FoodHnadlers.Get)           // Get food by ID
		foodGroup.GET("/list", a.FoodHnadlers.List)             // List all foods
		foodGroup.GET("/get/name/:name", a.FoodHnadlers.GetByName) // Get food by name
	}
}

