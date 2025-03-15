package entity

type Category struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) Category {
	return Category{
		Name: name,
	}
}
