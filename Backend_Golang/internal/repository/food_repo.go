package repository

import (
	"errors" // Add this import for error handling
	"fmt"    // Add this import for formatting errors
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
	if name == "" {
		return nil, errors.New("food name cannot be empty")
	}

	var food models.Food
	if err := r.db.Where("name = ?", name).First(&food).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if no record is found
		}
		return nil, fmt.Errorf("failed to find food by name: %w", err)
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
