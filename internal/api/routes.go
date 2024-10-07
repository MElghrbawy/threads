package api

import (
	"database/sql"
	"github.com/MElghrbawy/threads/internal/api/handler"
	"github.com/MElghrbawy/threads/internal/repository"
	"github.com/MElghrbawy/threads/internal/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	api := app.Group("/api")

	// Category routes
	categories := api.Group("/categories")
	categories.Post("", categoryHandler.CreateCategory)
	categories.Get("", categoryHandler.ListCategories)
	categories.Get("/:id", categoryHandler.GetCategory)
	categories.Put("", categoryHandler.UpdateCategory)
	categories.Delete("/:id", categoryHandler.DeleteCategory)

}
