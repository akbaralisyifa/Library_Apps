package service

import (
	"errors"
	"library/internal/features/categories"
	"log"
)

type CategoryServices struct {
	qry categories.Query
}

func NewCategoryServices(q categories.Query) categories.Service {
	return &CategoryServices{
		qry: q,
	}
}

func (cs *CategoryServices) AddCategory(newCategory categories.Categories) error {
	err := cs.qry.AddCategory(newCategory)
	if err != nil {
		log.Print("add category sql error:", err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func (cs *CategoryServices) GetAllCategory()([]categories.Categories, error){
	result, err := cs.qry.GetAllCategory()
	if err != nil {
		log.Print("Get All Category sql error :", err.Error())
		return []categories.Categories{}, err
	}

	return result, nil
}


func (cs *CategoryServices) UpdateCategory(categoryID uint, updateCategory categories.Categories) error {
	err := cs.qry.UpdateCategory(categoryID, updateCategory);

	if err != nil {
		log.Print("Update Category Error :", err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func (cs *CategoryServices) DeleteCategory(categoryID uint) error {
	err := cs.qry.DeleteCategory(categoryID);

	if err != nil {
		log.Print("delete category error :", err.Error())
		return errors.New("internal server error")
	}

	return nil
}