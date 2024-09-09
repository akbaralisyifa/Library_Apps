package service

import (
	"errors"
	"library/internal/features/recomendation"
	"log"
)

type RecommendServices struct {
	qry recomendation.Query
}

func NewRecommendServices(q recomendation.Query) recomendation.Service {
	return &RecommendServices{
		qry: q,
	}
}

// add Recommend book
func (rs *RecommendServices) AddRecommend(newRecommend recomendation.Recomendation) error {
	err := rs.qry.AddRecommend(newRecommend)

		if err != nil {
			log.Print("Add recommendation book query error")
			return errors.New("error in server")
		}

	return nil
}

// get All recommend book
func (rs *RecommendServices) GetAllRecommend() ([]recomendation.Recomendation, error){
	result, err := rs.qry.GetAllRecommend()

	if err != nil {
		log.Print("get all recommend book query error")
		return []recomendation.Recomendation{}, errors.New("error in server")
	}

	return result, nil
}

// update recommend book
func (rs *RecommendServices) UpdateRecommend(recommendID uint, updateRecommend recomendation.Recomendation) error {
	err := rs.qry.UpdateRecommend(recommendID, updateRecommend);

	if err != nil {
		log.Print("update recommend book query error")
		return errors.New("error in server")
	}

	return nil
}

// delete recommend book
func (rs *RecommendServices) DeleteRecommend(recommendID uint) error {
	err := rs.qry.DeleteRecommend(recommendID);

	if err != nil {
		log.Print("delete recommend book query error")
		return errors.New("error in server")
	}

	return nil
}