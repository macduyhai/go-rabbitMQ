package models

type Product struct {
	Name     string  `json:"name" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
}
