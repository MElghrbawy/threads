package repository

import (
	"context"
	"database/sql"
	"github.com/MElghrbawy/print/internal/db/sqlc"
	"github.com/MElghrbawy/print/internal/models"
)

type CategoryRepository struct {
	q *sqlc.Queries
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		q: sqlc.New(db),
	}
}

// GetCategory fetches a category by its ID
func (r *CategoryRepository) GetCategory(ctx context.Context, id int32) (*models.Category, error) {
	dbCategory, err := r.q.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.Category{
		ID:   dbCategory.ID,
		Name: dbCategory.Name,
	}, nil
}

// CreateCategory creates a new category
func (r *CategoryRepository) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	dbCategory, err := r.q.CreateCategory(ctx, category.Name)
	if err != nil {
		return nil, err
	}
	return &models.Category{
		ID:   dbCategory.ID,
		Name: dbCategory.Name,
	}, nil
}

// ListCategories fetches all categories
func (r *CategoryRepository) ListCategories(ctx context.Context) ([]*models.Category, error) {
	dbCategories, err := r.q.ListCategories(ctx)
	if err != nil {
		return nil, err
	}

	var categories []*models.Category
	for _, dbCategory := range dbCategories {
		categories = append(categories, &models.Category{
			ID:   dbCategory.ID,
			Name: dbCategory.Name,
		})
	}
	return categories, nil
}

// UpdateCategory updates a category by its ID
func (r *CategoryRepository) UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	params := sqlc.UpdateCategoryParams{
		Name: category.Name,
		ID:   category.ID,
	}
	err := r.q.UpdateCategory(ctx, params)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// DeleteCategory deletes a category by its ID
func (r *CategoryRepository) DeleteCategory(ctx context.Context, id int32) error {
	return r.q.DeleteCategory(ctx, id)
}
