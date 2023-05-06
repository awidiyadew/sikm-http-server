package model

type Product struct{
	ID string `json:"id"` // struct tag
	Name string `json:"name"`
	Price string `json:"price"`
	Quantity int `json:"qty"`
}