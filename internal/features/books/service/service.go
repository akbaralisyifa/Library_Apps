package service

import (
	"errors"
	"library/internal/features/books"
	"library/internal/features/users"
	"log"
)

type BookServices struct {
	qry books.Query
	usr users.Service
}

func NewBookServices(q books.Query, u users.Service) books.Service {
	return &BookServices{
		qry: q,
		usr: u,
	}
}

func (bs *BookServices) AddBook(userID uint, newBook books.Books) error {

	val, err := bs.usr.GetUser(userID);
	if err != nil {
		log.Print("Get User query error", err.Error())
		return errors.New("error in server")
	}

	if val.Role == "admin" {
        err = bs.qry.AddBook(newBook)
		if err != nil {
			log.Print("Add book query error")
			return errors.New("error in server")
		}
		return nil
	}
	
	if val.Role == "user" {
		return errors.New("failed: user does not have permission to add books")
	}

	return errors.New("failed: unknown user role")
}

func (bs *BookServices) GetAllBook() ([]books.Books, error) {
	result, err := bs.qry.GetAllBook()
	if err != nil {
		log.Print("get all book error", err.Error())
		return []books.Books{}, errors.New("server in error")
	}

	return result, nil
}

func (bs *BookServices) GetBook(bookID uint) (books.Books, error){
	result, err := bs.qry.GetBook(bookID)

	if err != nil {
		log.Print("get book by id error", err.Error())
		return books.Books{}, errors.New("error in server")
	}

	return result, nil
}

func (bs *BookServices) UpdateBook(userID, bookID uint, updateBook books.Books) error {

	val, err := bs.usr.GetUser(userID);
	if err != nil {
		log.Print("Get User query error", err.Error())
		return errors.New("error in server")
	}

	if val.Role == "admin" {
        err = bs.qry.UpdateBook(bookID, updateBook)
		if err != nil {
			log.Print("Update book query error")
			return errors.New("error in server")
		}
		return nil
	}

	if val.Role == "user" {
		return errors.New("failed: user does not have permission to add books")
	}

	return errors.New("failed: unknown user role")
}

func (bs *BookServices) DeleteBook(userID, bookID uint) error {

	val, err := bs.usr.GetUser(userID);
	if err != nil {
		log.Print("Get User query error", err.Error())
		return errors.New("error in server")
	}

	if val.Role == "admin" {
		err := bs.qry.DeleteBook(bookID)
		if err != nil {
			log.Print("Update book query error")
			return errors.New("error in server")
		}
		return nil
	}

	if val.Role == "user" {
		return errors.New("failed: user does not have permission to add books")
	}

	return errors.New("failed: unknown user role")
}