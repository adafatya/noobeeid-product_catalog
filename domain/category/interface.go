package category

import "github.com/adafatya/noobeeid-product_catalog/entity"

type CategoryRepositoryInterface interface {
	// get all category
	//
	// return list of category, error
	GetAll() ([]entity.Category, error)
}
