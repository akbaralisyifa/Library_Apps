package repository

import (
	"library/internal/features/categories"

	"gorm.io/gorm"
)

type CategoryModels struct {
	db gorm.DB
}

func NewCategoryModels(connect *gorm.DB) categories.Query{
	return &CategoryModels{
		db: *connect,
	}
}

// Create Category
func (cm *CategoryModels) AddCategory(newCategory categories.Categories) error {
	cnvData := ToCategoryQuery(newCategory)
	err := cm.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	return nil
}

// get All Category
func (cm *CategoryModels) GetAllCategory() ([]categories.Categories, error){
	var result []Categories
	var resultMap []categories.Categories
	err := cm.db.Find(&result).Error

	if err != nil {
		return []categories.Categories{}, err
	}

	for _, val := range result {
		resultMap = append(resultMap, val.ToCategoryEntity())
	}

	return resultMap, nil
}

// Update Category
func (cm *CategoryModels) UpdateCategory(categoryID uint, updateCategory categories.Categories) error {
	cnvData := ToCategoryQuery(updateCategory);
	qry := cm.db.Where("id = ?", categoryID).Updates(&cnvData);

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Delete Category
func (cm *CategoryModels) DeleteCategory(categoryID uint) error {
	qry := cm.db.Where("id = ?", categoryID).Delete(&Categories{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}