package service

import (
	"context"
	"github.com/MElghrbawy/print/internal/models"
	"github.com/MElghrbawy/print/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// GetCategory returns a category by ID
func (s *CategoryService) GetCategory(ctx context.Context, id int32) (*models.Category, error) {
	return s.repo.GetCategory(ctx, id)
}

// CreateCategory creates a new category
func (s *CategoryService) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	return s.repo.CreateCategory(ctx, category)
}

// ListCategories returns all categories
func (s *CategoryService) ListCategories(ctx context.Context) ([]*models.Category, error) {
	return s.repo.ListCategories(ctx)
}

// UpdateCategory updates an existing category
func (s *CategoryService) UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	return s.repo.UpdateCategory(ctx, category)
}

// DeleteCategory deletes a category by ID
func (s *CategoryService) DeleteCategory(ctx context.Context, id int32) error {
	return s.repo.DeleteCategory(ctx, id)
}
