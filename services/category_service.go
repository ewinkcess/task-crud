package services

import (
	"errors"
	"strings"
	"task-crud/models"
	"task-crud/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(category *models.Category) error {
	// Validasi
	if strings.TrimSpace(category.Name) == "" {
		return errors.New("name wajib diisi")
	}

	// Cegah duplicate name
	existing, err := s.repo.GetByName(category.Name)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("category dengan nama tersebut sudah ada")
	}

	return s.repo.Create(category)
}
