package handlers

import "library/internal/features/categories"

type categoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToCategoryResponse(result []categories.Categories) []categoryResponse {
	var categories []categoryResponse

	for _, val := range result {
		categories = append(categories, categoryResponse{
			ID:   val.ID,
			Name: val.Name,
		})
	}

	return	categories
}
