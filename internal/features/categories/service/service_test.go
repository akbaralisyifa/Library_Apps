package service_test

import (
	"errors"
	"library/internal/features/categories"
	"library/internal/features/categories/service"
	"library/internal/features/users"
	"library/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCategory(t *testing.T) {
	mockQry := mocks.NewQuery(t)   // Mock untuk categories.Query
	mockUsr := mocks.NewService(t) // Mock untuk users.Service
	categorySrv := service.NewCategoryServices(mockQry, mockUsr)

	t.Run("Success Add Category by Admin", func(t *testing.T) {
		newCategory := categories.Categories{Name: "Fiction"}
		userID := uint(1)

		// Mock GetUser untuk return user dengan role "admin"
		mockUsr.On("GetUser", userID).Return(users.Users{Role: "admin"}, nil).Once()
		mockQry.On("AddCategory", newCategory).Return(nil).Once()

		err := categorySrv.AddCategory(userID, newCategory)

		assert.NoError(t, err)
		mockUsr.AssertExpectations(t)
		mockQry.AssertExpectations(t)
	})

	t.Run("Fail Add Category by Non-Admin", func(t *testing.T) {
		newCategory := categories.Categories{Name: "Fiction"}
		userID := uint(2)

		// Mock GetUser untuk return user dengan role "user"
		mockUsr.On("GetUser", userID).Return(users.Users{Role: "user"}, nil).Once()

		err := categorySrv.AddCategory(userID, newCategory)

		assert.Error(t, err)
		assert.Equal(t, "failed: user does not have permission to add books", err.Error())
		mockUsr.AssertExpectations(t)
	})

	t.Run("Error GetUser", func(t *testing.T) {
		newCategory := categories.Categories{Name: "Fiction"}
		userID := uint(3)

		// Mock GetUser dengan return error
		mockUsr.On("GetUser", userID).Return(users.Users{}, errors.New("query error")).Once()

		err := categorySrv.AddCategory(userID, newCategory)

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		mockUsr.AssertExpectations(t)
	})
}

func TestGetAllCategory(t *testing.T) {
	mockQry := mocks.NewQuery(t)
	mockUsr := mocks.NewService(t)
	categorySrv := service.NewCategoryServices(mockQry, mockUsr)

	t.Run("Success Get All Categories", func(t *testing.T) {
		expectedCategories := []categories.Categories{
			{Name: "Fiction"},
			{Name: "Science"},
		}

		mockQry.On("GetAllCategory").Return(expectedCategories, nil).Once()

		result, err := categorySrv.GetAllCategory()

		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, result)
		mockQry.AssertExpectations(t)
	})

	t.Run("Error Get All Categories", func(t *testing.T) {
		mockQry.On("GetAllCategory").Return([]categories.Categories{}, errors.New("query error")).Once()

		result, err := categorySrv.GetAllCategory()

		assert.Error(t, err)
		assert.Equal(t, "query error", err.Error())
		assert.Empty(t, result)
		mockQry.AssertExpectations(t)
	})
}

func TestUpdateCategory(t *testing.T) {
	mockQry := mocks.NewQuery(t)
	mockUsr := mocks.NewService(t)
	categorySrv := service.NewCategoryServices(mockQry, mockUsr)

	t.Run("Success Update Category by Admin", func(t *testing.T) {
		updateCategory := categories.Categories{Name: "Updated Fiction"}
		userID := uint(1)
		categoryID := uint(1)

		mockUsr.On("GetUser", userID).Return(users.Users{Role: "admin"}, nil).Once()
		mockQry.On("UpdateCategory", categoryID, updateCategory).Return(nil).Once()

		err := categorySrv.UpdateCategory(userID, categoryID, updateCategory)

		assert.NoError(t, err)
		mockUsr.AssertExpectations(t)
		mockQry.AssertExpectations(t)
	})

	t.Run("Fail Update Category by Non-Admin", func(t *testing.T) {
		updateCategory := categories.Categories{Name: "Updated Fiction"}
		userID := uint(2)
		categoryID := uint(1)

		mockUsr.On("GetUser", userID).Return(users.Users{Role: "user"}, nil).Once()

		err := categorySrv.UpdateCategory(userID, categoryID, updateCategory)

		assert.Error(t, err)
		assert.Equal(t, "failed: user does not have permission to add books", err.Error())
		mockUsr.AssertExpectations(t)
	})

	t.Run("Error GetUser", func(t *testing.T) {
		updateCategory := categories.Categories{Name: "Updated Fiction"}
		userID := uint(3)
		categoryID := uint(1)

		mockUsr.On("GetUser", userID).Return(users.Users{}, errors.New("query error")).Once()

		err := categorySrv.UpdateCategory(userID, categoryID, updateCategory)

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		mockUsr.AssertExpectations(t)
	})
}

func TestDeleteCategory(t *testing.T) {
	mockQry := mocks.NewQuery(t)
	mockUsr := mocks.NewService(t)
	categorySrv := service.NewCategoryServices(mockQry, mockUsr)

	t.Run("Success Delete Category by Admin", func(t *testing.T) {
		userID := uint(1)
		categoryID := uint(1)

		mockUsr.On("GetUser", userID).Return(users.Users{Role: "admin"}, nil).Once()
		mockQry.On("DeleteCategory", categoryID).Return(nil).Once()

		err := categorySrv.DeleteCategory(userID, categoryID)

		assert.NoError(t, err)
		mockUsr.AssertExpectations(t)
		mockQry.AssertExpectations(t)
	})

	t.Run("Fail Delete Category by Non-Admin", func(t *testing.T) {
		userID := uint(2)
		categoryID := uint(1)

		mockUsr.On("GetUser", userID).Return(users.Users{Role: "user"}, nil).Once()

		err := categorySrv.DeleteCategory(userID, categoryID)

		assert.Error(t, err)
		assert.Equal(t, "failed: user does not have permission to add books", err.Error())
		mockUsr.AssertExpectations(t)
	})

	t.Run("Error GetUser", func(t *testing.T) {
		userID := uint(3)
		categoryID := uint(1)

		mockUsr.On("GetUser", userID).Return(users.Users{}, errors.New("query error")).Once()

		err := categorySrv.DeleteCategory(userID, categoryID)

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		mockUsr.AssertExpectations(t)
	})
}
