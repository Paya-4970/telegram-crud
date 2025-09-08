package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Paya-4970/telegram-crud/internal/app"
)

func InitRouter(e *gin.Engine, a *app.App) {
	v1 := e.Group("food")
	{
		v1.GET("/get/:id", a.FoodHnadlers.Get)
		v1.GET("/list", a.FoodHnadlers.List)
		v1.GET("/get/name/:name", a.FoodHnadlers.GetByName)
	}
}
