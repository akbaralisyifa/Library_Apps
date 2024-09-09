package handlers

import "library/internal/features/recomendation"

type RecommnedInput struct {
	Reason string `json:"reason"`
}

func ToModelRecommend(userID, bookID uint, ri RecommnedInput) recomendation.Recomendation {
	return recomendation.Recomendation{
		UserID: userID,
		BookID: bookID,
		Reason: ri.Reason,
	}
}

func ToModelRecommendUpdate(ri RecommnedInput) recomendation.Recomendation {
	return recomendation.Recomendation{
		Reason: ri.Reason,
	}
}