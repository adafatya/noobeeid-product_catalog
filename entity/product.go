package entity

type Product struct {
	Id         uint   `json:"id"`
	MerchantId uint   `json:"merchant_id"`
	Name       string `json:"name"`
	Stock      uint   `json:"stock"`
	Price      uint   `json:"price"`
	CategoryId uint   `json:"category_id"`

	Category Category `json:"category"`
}

func NewProduct(
	merchant_id uint,
	name string,
	stock uint,
	price uint,
	categoryId uint,
) Product {
	return Product{
		MerchantId: merchant_id,
		Name:       name,
		Stock:      stock,
		Price:      price,
		CategoryId: categoryId,
	}
}
