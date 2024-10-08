// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	books "library/internal/features/books"
	"library/internal/features/categories"
	"library/internal/features/recomendation"
	"library/internal/features/users"

	mock "github.com/stretchr/testify/mock"
)

// Query is an autogenerated mock type for the Query type
type Query struct {
	mock.Mock
}

func (m *Query) Register(user users.Users) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *Query) GetUser(userID uint) (users.Users, error) {
    args := m.Called(userID)
    return args.Get(0).(users.Users), args.Error(1)
}

func (m *Query) UpdateUser(userID uint, user users.Users) error {
    args := m.Called(userID, user)
    return args.Error(0)
}

func (m *Query) DeleteUser(userID uint) error {
    args := m.Called(userID)
    return args.Error(0)
}

func (m *Query) Login(email string) (users.Users, error) {
    args := m.Called(email)
    return args.Get(0).(users.Users), args.Error(1)
}

// AddBook provides a mock function with given fields: newBook
func (_m *Query) AddBook(newBook books.Books) error {
	ret := _m.Called(newBook)

	if len(ret) == 0 {
		panic("no return value specified for AddBook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(books.Books) error); ok {
		r0 = rf(newBook)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBook provides a mock function with given fields: bookID
func (_m *Query) DeleteBook(bookID uint) error {
	ret := _m.Called(bookID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(bookID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBook provides a mock function with given fields: title
func (_m *Query) GetAllBook(title string) ([]books.Books, error) {
	ret := _m.Called(title)

	if len(ret) == 0 {
		panic("no return value specified for GetAllBook")
	}

	var r0 []books.Books
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]books.Books, error)); ok {
		return rf(title)
	}
	if rf, ok := ret.Get(0).(func(string) []books.Books); ok {
		r0 = rf(title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]books.Books)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBook provides a mock function with given fields: bookID
func (_m *Query) GetBook(bookID uint) (books.Books, error) {
	ret := _m.Called(bookID)

	if len(ret) == 0 {
		panic("no return value specified for GetBook")
	}

	var r0 books.Books
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (books.Books, error)); ok {
		return rf(bookID)
	}
	if rf, ok := ret.Get(0).(func(uint) books.Books); ok {
		r0 = rf(bookID)
	} else {
		r0 = ret.Get(0).(books.Books)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBook provides a mock function with given fields: bookID, updateBook
func (_m *Query) UpdateBook(bookID uint, updateBook books.Books) error {
	ret := _m.Called(bookID, updateBook)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, books.Books) error); ok {
		r0 = rf(bookID, updateBook)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *Query) AddCategory(newCategory categories.Categories) error {
	args := m.Called(newCategory)
	return args.Error(0)
}

func (m *Query) GetAllCategory() ([]categories.Categories, error) {
	args := m.Called()
	return args.Get(0).([]categories.Categories), args.Error(1)
}

func (m *Query) UpdateCategory(categoryID uint, updateCategory categories.Categories) error {
	args := m.Called(categoryID, updateCategory)
	return args.Error(0)
}

func (m *Query) DeleteCategory(categoryID uint) error {
	args := m.Called(categoryID)
	return args.Error(0)
}

func (m *Query) AddRecommend(recommendation recomendation.Recomendation) error {
    args := m.Called(recommendation)
    return args.Error(0)
}

func (m *Query) GetAllRecommend() ([]recomendation.Recomendation, error) {
    args := m.Called()
    return args.Get(0).([]recomendation.Recomendation), args.Error(1)
}

func (m *Query) UpdateRecommend(id uint, recommendation recomendation.Recomendation) error {
    args := m.Called(id, recommendation)
    return args.Error(0)
}

func (m *Query) DeleteRecommend(id uint) error {
    args := m.Called(id)
    return args.Error(0)
}

// NewQuery creates a new instance of Query. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *Query {
	mock := &Query{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
