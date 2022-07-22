package models

type UserBuy struct {
	Username string `json:"username" binding:"required"`
	Products []Product
}
