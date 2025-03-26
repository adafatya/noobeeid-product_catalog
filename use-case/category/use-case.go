package category

import (
	"github.com/adafatya/noobeeid-product_catalog/constant"
	"github.com/adafatya/noobeeid-product_catalog/entity"
)

type CategoryUseCase struct {
	repo CategoryRepositoryInterface
}

func NewCategoryUseCase(repo CategoryRepositoryInterface) CategoryUseCase {
	return CategoryUseCase{
		repo: repo,
	}
}

func (uc CategoryUseCase) GetAllCategory() ([]entity.Category, *constant.Error) {
	// get data from db
	categories, err := uc.repo.GetAll()
	if err != nil {
		return nil, constant.ErrRepositoryError
	}

	return categories, nil
}
