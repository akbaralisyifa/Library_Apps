package repository

import (
	"library/internal/features/recomendation"

	"gorm.io/gorm"
)

type RecommendModels struct {
	db *gorm.DB
}

func NewRecommendModels(connect *gorm.DB) recomendation.Query {
	return &RecommendModels{
		db: connect,
	}
}


func (rm *RecommendModels) AddRecommend(newRecommend recomendation.Recomendation) error {
	cnvData := ToRecommendQuery(newRecommend)
	err := rm.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	return nil
}

func (rm *RecommendModels) GetAllRecommend() ([]recomendation.Recomendation, error) {
	var result []Recomendation
	var resultMap []recomendation.Recomendation
	err := rm.db.Find(&result).Error

	if err != nil {
		return []recomendation.Recomendation{}, err
	}

	for _, val := range result{
		resultMap = append(resultMap, val.ToRecommendEntity())
	}

	return resultMap, nil
}


func (rm *RecommendModels) UpdateRecommend(recommendID uint, updateRecommend recomendation.Recomendation) error {
	cnvData := ToRecommendQuery(updateRecommend)
	qry := rm.db.Where("id = ?", recommendID).Updates(&cnvData);

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (rm *RecommendModels) DeleteRecommend(recommendID uint) error {
	qry := rm.db.Where("id = ?", recommendID).Delete(&Recomendation{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}