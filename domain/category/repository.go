package category

import (
	"database/sql"
	"log"

	"github.com/adafatya/noobeeid-product_catalog/entity"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return CategoryRepository{
		db: db,
	}
}

func (r CategoryRepository) GetAll() ([]entity.Category, error) {
	categories := make([]entity.Category, 0)

	err := r.db.Select(&categories, "SELECT id, name FROM categories")
	if err != nil && err != sql.ErrNoRows {
		log.Println("terdapat error saat mengambil category:", err)
		return nil, err
	}

	return categories, nil
}
