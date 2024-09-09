package handlers

import "library/internal/features/recomendation"

type RecommnedResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	BookID uint   `json:"book_id"`
	Reason string `json:"reason"`
}

func ToRecommendResponse(result []recomendation.Recomendation) []RecommnedResponse{
	var responseData []RecommnedResponse

	for _, val := range result{
		responseData = append(responseData, RecommnedResponse{
			ID: val.ID,
			UserID: val.UserID,
			BookID: val.BookID,
			Reason: val.Reason,
		})
	}

	return responseData
}