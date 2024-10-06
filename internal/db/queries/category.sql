-- name: GetCategoryByID :one
SELECT id, name, created_at, updated_at
FROM category
WHERE id = $1;
-- name: ListCategories :many
SELECT id, name, created_at, updated_at
FROM category
ORDER BY created_at DESC;
-- name: CreateCategory :one
INSERT INTO category (name)
VALUES ($1)
RETURNING id, name, created_at, updated_at;
-- name: UpdateCategory :exec
UPDATE category
SET name       = $1,
    updated_at = NOW()
WHERE id = $2;
-- name: DeleteCategory :exec
DELETE
FROM category
WHERE id = $1;
