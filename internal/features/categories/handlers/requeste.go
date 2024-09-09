package handlers

import "library/internal/features/categories"

type categoryInput struct {
	Name string `json:"name"`
}

func ToModelCategory(ci categoryInput) categories.Categories {
	return categories.Categories{
		Name: ci.Name,
	}
}