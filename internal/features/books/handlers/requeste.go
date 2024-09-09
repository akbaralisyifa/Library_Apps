package handlers

import "library/internal/features/books"

type BookInput struct {
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear string `json:"published_year"`
}

func ToModelBook(bi BookInput) books.Books{
	return books.Books{
		CategoryID:	   bi.CategoryID,
		Title:		   bi.Title,
		Author: 	   bi.Author,
		PublishedYear: bi.PublishedYear,
	}
}