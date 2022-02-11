package model

import (
	"encoding/json"
	"time"
)

type Product struct {
	ID          int
	Title       string
	Price       int
	Description string
	Rating      int
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	Discount    json.Number `json:"discount" binding:"required,number"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	Discount    int    `json:"discount"`
}
