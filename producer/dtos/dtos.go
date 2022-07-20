package dtos

import "github.com/macduyhai/go-rabbitMQ/producer/models"

type Response struct {
	Data interface{} `json:"data"`
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
}

type BuyRequest struct {
	Username string           `json:"username" binding:"required"`
	Token    string           `json:"token" binding:"required"`
	Products []models.Product `json:"products" binding:"required,dive"`
}

type DeleteRequest struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
