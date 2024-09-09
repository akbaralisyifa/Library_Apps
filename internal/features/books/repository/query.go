package repository

import (
	"library/internal/features/books"

	"gorm.io/gorm"
)

type BookModels struct {
	db *gorm.DB
}

func NewBookModels(connect *gorm.DB) books.Query {
	return &BookModels{
		db: connect,
	}
}

// Create book
func (bm *BookModels) AddBook(newBook books.Books) error {
	cnvData := ToBookQuery(newBook)

	err := bm.db.Create(&cnvData).Error
	if err != nil {
		return err
	}

	return nil
}

// get Book
func (bm *BookModels) GetAllBook() ([]books.Books, error){
	var result []Books
	var resultMap []books.Books
	err := bm.db.Find(&result).Error
	if err != nil {
		return []books.Books{}, err
	}

	for _, val := range result {
		resultMap = append(resultMap, val.ToBookEntity())
	}

	return resultMap, nil
}

func (bm *BookModels) GetBook(bookID uint) (books.Books, error) {
	var result Books
	err := bm.db.Where("id = ?", bookID).First(&result).Error
	if err != nil {
		return books.Books{}, err
	}

	return result.ToBookEntity(), nil
}

// Update book
func (bm *BookModels) UpdateBook(bookID uint, updateBook books.Books) error {
	cnvData := ToBookQuery(updateBook);

	qry := bm.db.Where("id = ?", bookID).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}


// Delete book
func (bm *BookModels) DeleteBook(bookID uint) error {
	qry := bm.db.Where("id = ?", bookID).Delete(&Books{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}