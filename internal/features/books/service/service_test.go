package service_test

import (
	"errors"
	"library/internal/features/books"
	"library/internal/features/books/service"
	"library/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
    qry := mocks.NewQuery(t)
    usrSrv := mocks.NewService(t)
    bookSrv := service.NewBookServices(qry, usrSrv)

    t.Run("Success Add Book by Admin", func(t *testing.T) {
        userID := uint(1)
        newBook := books.Books{
            Title:  "Go Programming",
            Author: "John Doe",
        }

        // Mock GetUser untuk return user dengan role "admin"
        usrSrv.On("GetUser", userID).Return(&books.User{ID: userID, Role: "admin"}, nil).Once()
        
        // Pastikan bahwa AddBook di query dipanggil dengan argumen yang sesuai
        qry.On("AddBook", newBook).Return(nil).Once()

        err := bookSrv.AddBook(userID, newBook)

        assert.NoError(t, err)
        usrSrv.AssertExpectations(t)
        qry.AssertExpectations(t)
    })

    t.Run("Fail Add Book by Non-Admin", func(t *testing.T) {
        userID := uint(1)
        newBook := books.Books{
            Title:  "Go Programming",
            Author: "John Doe",
        }

        // Mock GetUser untuk return user dengan role bukan "admin"
        usrSrv.On("GetUser", userID).Return(&books.User{ID: userID, Role: "user"}, nil).Once()

        err := bookSrv.AddBook(userID, newBook)

        assert.Error(t, err)
        assert.Equal(t, "failed: user does not have permission to add books", err.Error())
        usrSrv.AssertExpectations(t)
        qry.AssertExpectations(t)
    })

    t.Run("Fail Add Book with Error from Query", func(t *testing.T) {
        userID := uint(1)
        newBook := books.Books{
            Title:  "Go Programming",
            Author: "John Doe",
        }

        // Mock GetUser untuk return user dengan role "admin"
        usrSrv.On("GetUser", userID).Return(&books.User{ID: userID, Role: "admin"}, nil).Once()
        
        // Simulasi error saat menambahkan buku
        qry.On("AddBook", newBook).Return(errors.New("database error")).Once()

        err := bookSrv.AddBook(userID, newBook)

        assert.Error(t, err)
        assert.Equal(t, "database error", err.Error())
        usrSrv.AssertExpectations(t)
        qry.AssertExpectations(t)
    })
}

func TestGetAllBooks(t *testing.T) {
	qry := mocks.NewQuery(t)
	usrSrv := mocks.NewService(t)
	bookSrv := service.NewBookServices(qry, usrSrv)

	t.Run("Success Get All Books", func(t *testing.T) {
		title := "Go"
		expectedResult := []books.Books{
			{ID: 1, Title: "Go Programming", Author: "John Doe"},
		}

		qry.On("GetAllBook", title).Return(expectedResult, nil).Once()

		result, err := bookSrv.GetAllBook(title)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
		qry.AssertExpectations(t)
	})

	t.Run("Error Get All Books", func(t *testing.T) {
		title := "Go"

		qry.On("GetAllBook", title).Return([]books.Books{}, errors.New("error fetching books")).Once()

		result, err := bookSrv.GetAllBook(title)

		assert.Error(t, err)
		assert.Equal(t, "server in error", err.Error())
		assert.Empty(t, result)
		qry.AssertExpectations(t)
	})
}

func TestGetBook(t *testing.T) {
	qry := mocks.NewQuery(t)
	usrSrv := mocks.NewService(t)
	bookSrv := service.NewBookServices(qry, usrSrv)

	t.Run("Success Get Book", func(t *testing.T) {
		bookID := uint(1)
		expectedBook := books.Books{
			ID:      bookID,
			Title:   "Go Programming",
			Author:  "John Doe",
		}

		qry.On("GetBook", bookID).Return(expectedBook, nil).Once()

		result, err := bookSrv.GetBook(bookID)

		assert.NoError(t, err)
		assert.Equal(t, expectedBook, result)
		qry.AssertExpectations(t)
	})

	t.Run("Error Get Book", func(t *testing.T) {
		bookID := uint(1)

		qry.On("GetBook", bookID).Return(books.Books{}, errors.New("error fetching book")).Once()

		result, err := bookSrv.GetBook(bookID)

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		assert.Empty(t, result)
		qry.AssertExpectations(t)
	})
}

func TestUpdateBook(t *testing.T) {
	qry := mocks.NewQuery(t)
	usrSrv := mocks.NewService(t)
	bookSrv := service.NewBookServices(qry, usrSrv)

	t.Run("Success Update Book by Admin", func(t *testing.T) {
		userID := uint(1)
		bookID := uint(1)
		updateBook := books.Books{Title: "Advanced Go Programming"}

		// Mock GetUser untuk return user dengan role "admin"
		usrSrv.On("GetUser", userID).Return(&books.User{ID: userID, Role: "admin"}, nil).Once()
		qry.On("UpdateBook", bookID, updateBook).Return(nil).Once()

		err := bookSrv.UpdateBook(userID, bookID, updateBook)

		assert.NoError(t, err)
		usrSrv.AssertExpectations(t)
		qry.AssertExpectations(t)
	})

	t.Run("Error Update Book by User", func(t *testing.T) {
		userID := uint(2)
		bookID := uint(1)
		updateBook := books.Books{Title: "Advanced Go Programming"}

		// Mock GetUser untuk return user dengan role "user"
		usrSrv.On("GetUser", userID).Return(&books.User{ID: userID, Role: "user"}, nil).Once()

		err := bookSrv.UpdateBook(userID, bookID, updateBook)

		assert.Error(t, err)
		assert.Equal(t, "failed: user does not have permission to update books", err.Error())
		usrSrv.AssertExpectations(t)
	})

	t.Run("Error GetUser Failure", func(t *testing.T) {
		userID := uint(1)
		bookID := uint(1)
		updateBook := books.Books{Title: "Advanced Go Programming"}

		// Mock GetUser dengan return error
		usrSrv.On("GetUser", userID).Return(nil, errors.New("error fetching user")).Once()

		err := bookSrv.UpdateBook(userID, bookID, updateBook)

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		usrSrv.AssertExpectations(t)
	})
}

func TestDeleteBook(t *testing.T) {
	qry := mocks.NewQuery(t)
	usrSrv := mocks.NewService(t)
	bookSrv := service.NewBookServices(qry, usrSrv)

	t.Run("Success Delete Book by Admin", func(t *testing.T) {
		userID := uint(1)
		bookID := uint(1)

		// Mock GetUser untuk return user dengan role "admin"
		usrSrv.On("GetUser", userID).Return(&books.User{ID: userID, Role: "admin"}, nil).Once()
		qry.On("DeleteBook", bookID).Return(nil).Once()

		err := bookSrv.DeleteBook(userID, bookID)

		assert.NoError(t, err)
		usrSrv.AssertExpectations(t)
		qry.AssertExpectations(t)
	})

	t.Run("Error Delete Book by User", func(t *testing.T) {
		userID := uint(2)
		bookID := uint(1)

		// Mock GetUser untuk return user dengan role "user"
		usrSrv.On("GetUser", userID).Return(&books.User{ID: userID, Role: "user"}, nil).Once()

		err := bookSrv.DeleteBook(userID, bookID)

		assert.Error(t, err)
		assert.Equal(t, "failed: user does not have permission to delete books", err.Error())
		usrSrv.AssertExpectations(t)
	})

	t.Run("Error GetUser Failure", func(t *testing.T) {
		userID := uint(1)
		bookID := uint(1)

		// Mock GetUser dengan return error
		usrSrv.On("GetUser", userID).Return(nil, errors.New("error fetching user")).Once()

		err := bookSrv.DeleteBook(userID, bookID)

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		usrSrv.AssertExpectations(t)
	})
}
