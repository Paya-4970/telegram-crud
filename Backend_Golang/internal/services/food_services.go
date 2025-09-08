package services

import (
	"github.com/Paya-4970/telegram-crud/internal/models"
	"github.com/Paya-4970/telegram-crud/internal/repository"
)

type FoodServices interface {
	GetByID(id uint) (*models.Food, error)
	ListFood() ([]models.Food, error)
	GetFoodByName(name string) (*models.Food, error)
}

type foodServices struct {
	repo repository.FoodRepository
}

func NewFoodServices(repo repository.FoodRepository) FoodServices {
	return &foodServices{repo:repo}
}

func (s *foodServices) GetFoodByName(name string) (*models.Food, error) {
	return s.repo.FindByName(name)
}

func (s *foodServices) GetByID(id uint) (*models.Food, error) {
	return s.repo.FindByID(id)
}

func (s *foodServices) ListFood() ([]models.Food, error) {
	return s.repo.List()
}

