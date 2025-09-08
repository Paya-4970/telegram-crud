package repository

import (
	"github.com/Paya-4970/telegram-crud/internal/models"
	"gorm.io/gorm"
)

type FoodRepository interface {
	FindByID(id uint) (*models.Food, error)
	FindByName(name string) (*models.Food, error)
	List() ([]models.Food, error)
}

type foodRepo struct {
	db *gorm.DB
}

func NewFoodRepo(db *gorm.DB) FoodRepository {
	return &foodRepo{db: db}
}

func (r *foodRepo) FindByName(name string) (*models.Food, error) {
	var food models.Food
	if err := r.db.Where("name = ?", name).First(&food).Error; err != nil {
		return nil, err
	}
	return &food, nil
}

func (r *foodRepo) FindByID(id uint) (*models.Food, error) {
	var food models.Food
	if err := r.db.Where("id = ?", id).First(&food).Error; err != nil {
		return nil, err
	}
	return &food, nil
}

func (r *foodRepo) List() ([]models.Food, error) {
	var foods []models.Food
	if err := r.db.Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}
