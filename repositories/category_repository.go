package repositories

import (
	"database/sql"
	"task-crud/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) Create(category *models.Category) error {
	query := `
	INSERT INTO categories (name, description)
	VALUES ($1, $2)
	RETURNING id
	`

	return repo.db.QueryRow(
		query,
		category.Name,
		category.Description,
	).Scan(&category.ID)
}

func (repo *CategoryRepository) Exists(id int) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM categories WHERE id = $1)`
	var exists bool
	err := repo.db.QueryRow(query, id).Scan(&exists)
	return exists, err
}

func (repo *CategoryRepository) GetByName(name string) (*models.Category, error) {
	query := `SELECT id, name, description FROM categories WHERE name = $1`
	var c models.Category
	err := repo.db.QueryRow(query, name).Scan(&c.ID, &c.Name, &c.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &c, err
}
