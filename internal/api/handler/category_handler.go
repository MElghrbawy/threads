package handler

import (
	"context"
	"github.com/MElghrbawy/print/internal/models"
	"github.com/MElghrbawy/print/internal/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	svc *service.CategoryService
}

func NewCategoryHandler(svc *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: svc}
}

// GetCategory retrieves a category by ID
func (h *CategoryHandler) GetCategory(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid category ID")
	}

	category, err := h.svc.GetCategory(context.Background(), int32(id))
	if err != nil {
		return c.Status(http.StatusNotFound).SendString(err.Error())
	}

	return c.JSON(category)
}

// CreateCategory creates a new category
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	newCategory, err := h.svc.CreateCategory(context.Background(), &category)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(newCategory)
}

// ListCategories retrieves all categories
func (h *CategoryHandler) ListCategories(c *fiber.Ctx) error {
	categories, err := h.svc.ListCategories(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(categories)
}

// UpdateCategory updates an existing category
func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	// Ensure the category ID is valid
	if category.ID <= 0 {
		return c.Status(http.StatusBadRequest).SendString("Category ID is required")
	}

	// Call the repository method to update the category
	updatedCategory, err := h.svc.UpdateCategory(context.Background(), &category)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Return the updated category as a response
	return c.JSON(updatedCategory)
}

// DeleteCategory deletes a category by ID
func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid category ID")
	}

	if err := h.svc.DeleteCategory(context.Background(), int32(id)); err != nil {
		return c.Status(http.StatusNotFound).SendString(err.Error())
	}

	return c.SendStatus(http.StatusNoContent) // No content response for successful delete
}
