package handlers

import (
	"net/http"
	"strconv"

	"github.com/Paya-4970/telegram-crud/internal/services"
	"github.com/gin-gonic/gin"
)

type FoodHandlers struct {
	services services.FoodServices
}

func NewFoodServies(s services.FoodServices) *FoodHandlers {
	return &FoodHandlers{services: s}
}

func (h *FoodHandlers) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	food, err := h.services.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.JSON(200, food)
}

func (h *FoodHandlers) GetByName(c *gin.Context) {
	name := c.Param("name")
	if name == " "{
		c.JSON(http.StatusBadRequest, gin.H{"error":"name could not be empty"})
		return
	}
	food, err := h.services.GetFoodByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	} 
	c.JSON(200, food)
}

func (h *FoodHandlers) List(c *gin.Context) {
	foods, err := h.services.ListFood()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.JSON(200, foods)
}