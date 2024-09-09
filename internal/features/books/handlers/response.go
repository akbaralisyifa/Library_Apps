package handlers

import "library/internal/features/books"

type BookResponse struct {
	ID			  uint   `json:"id"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear string `json:"published_year"`
}

func ToBookResponse(result []books.Books) []BookResponse {
	var responseData []BookResponse

	for _, val := range result {
		responseData = append(responseData, ToBookResponseById(val))
	}

	return responseData
}

func ToBookResponseById(result books.Books) BookResponse{
	return BookResponse{
		ID:      	   result.ID,
		CategoryID:    result.CategoryID,
		Title: 		   result.Title,
		Author: 	   result.Author,
		PublishedYear: result.PublishedYear,
	}
}