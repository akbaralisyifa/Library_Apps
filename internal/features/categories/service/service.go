package service

import (
	"errors"
	"library/internal/features/categories"
	"library/internal/features/users"
	"log"
)

type CategoryServices struct {
	qry categories.Query
	usr users.Service
}

func NewCategoryServices(q categories.Query, s users.Service) categories.Service {
	return &CategoryServices{
		qry: q,
		usr: s,
	}
}

func (cs *CategoryServices) AddCategory(userID uint, newCategory categories.Categories) error {

	val, err := cs.usr.GetUser(userID);
	if err != nil {
		log.Print("Get User query error", err.Error())
		return errors.New("error in server")
	}

	if val.Role == "admin" {
		err = cs.qry.AddCategory(newCategory)
	if err != nil {
		log.Print("add category sql error:", err.Error())
		return errors.New("internal server error")
	}

	return nil
	}

	if val.Role == "user" {
		return errors.New("failed: user does not have permission to add books")
	}

	return errors.New("failed: unknown user role")
	
}

func (cs *CategoryServices) GetAllCategory()([]categories.Categories, error){
	result, err := cs.qry.GetAllCategory()
	if err != nil {
		log.Print("Get All Category sql error :", err.Error())
		return []categories.Categories{}, err
	}

	return result, nil
}


func (cs *CategoryServices) UpdateCategory(userID, categoryID uint, updateCategory categories.Categories) error {
	val, err := cs.usr.GetUser(userID);
	if err != nil {
		log.Print("Get User query error", err.Error())
		return errors.New("error in server")
	}

	if val.Role == "admin" {
		err := cs.qry.UpdateCategory(categoryID, updateCategory);
		if err != nil {
			log.Print("Update Category Error :", err.Error())
			return errors.New("internal server error")
		}
		return nil
	}

	if val.Role == "user" {
		return errors.New("failed: user does not have permission to add books")
	}

	return errors.New("failed: unknown user role")
}

func (cs *CategoryServices) DeleteCategory(userID, categoryID uint) error {
	val, err := cs.usr.GetUser(userID);
	if err != nil {
		log.Print("Get User query error", err.Error())
		return errors.New("error in server")
	}

	if val.Role == "admin" {
		err := cs.qry.DeleteCategory(categoryID);
		if err != nil {
			log.Print("delete category error :", err.Error())
			return errors.New("internal server error")
		}
		return nil
	}
	
	if val.Role == "user" {
		return errors.New("failed: user does not have permission to add books")
	}

	return errors.New("failed: unknown user role")
}