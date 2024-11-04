package dto

import "github.com/inventori-app-jeff/internal/model"

type CategoryResponse struct {
	ID           string        `json:"id"`
	Product      model.Product `json:"product"`
	CategoryName string        `json:"categoryName"`
}
