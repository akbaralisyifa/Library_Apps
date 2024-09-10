package service_test

import (
	"errors"
	"library/internal/features/users"
	"library/internal/features/users/service"
	"library/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	qry := mocks.NewQuery(t)  // Mock untuk users.Query
	jwt := mocks.NewJwtUtilityInteface(t) // Mock untuk utils.JwtUtilityInterface
	pwd := mocks.NewPasswordUtilityInterface(t) // Mock untuk utils.PasswordUtilityInterface
	usrSrv := service.NewUserServices(qry, jwt, pwd)

	t.Run("Success Register", func(t *testing.T) {
		newUser := users.Users{
			Username: "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		}

		// Mocking password hash generation
		pwd.On("GeneretePassword", "password123").Return([]byte("hashedpassword"), nil).Once()
		// Mocking the query to register user
		qry.On("Register", mock.Anything).Return(nil).Once()

		err := usrSrv.Register(newUser)

		assert.NoError(t, err)
		pwd.AssertExpectations(t)
		qry.AssertExpectations(t)
	})

	t.Run("Error Hashing Password", func(t *testing.T) {
		newUser := users.Users{
			Username:     "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		}

		// Mocking password hash generation to fail
		pwd.On("GeneretePassword", "password123").Return([]byte{}, errors.New("hash error")).Once()

		err := usrSrv.Register(newUser)

		assert.Error(t, err)
		assert.Equal(t, "hash error", err.Error())
		pwd.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
    qry := mocks.NewQuery(t)  // Mock untuk users.Query
    jwt := mocks.NewJwtUtilityInteface(t) // Mock untuk utils.JwtUtilityInterface
    pwd := mocks.NewPasswordUtilityInterface(t) // Mock untuk utils.PasswordUtilityInterface
    usrSrv := service.NewUserServices(qry, jwt, pwd)

    t.Run("Success Login", func(t *testing.T) {
        email := "john@example.com"
        password := "password123"
        userData := users.Users{
            ID:       1,
            Email:    "john@example.com",
            Password: "hashedpassword",
        }

        // Mocking Login query and password checking
        qry.On("Login", email).Return(userData, nil).Once()
        pwd.On("CheckPassword", []byte(password), []byte("hashedpassword")).Return(nil).Once()
        jwt.On("GenerateJwt", userData.ID).Return("token123", nil).Once()

        result, token, err := usrSrv.Login(email, password)

        assert.NoError(t, err)
        assert.Equal(t, userData, result)
        assert.Equal(t, "token123", token)
        qry.AssertExpectations(t)
        pwd.AssertExpectations(t)
        jwt.AssertExpectations(t)
    })

    t.Run("Error Invalid Password", func(t *testing.T) {
        email := "john@example.com"
        password := "wrongpassword"
        userData := users.Users{
            ID:       1,
            Email:    "john@example.com",
            Password: "hashedpassword",
        }

        qry.On("Login", email).Return(userData, nil).Once()
        pwd.On("CheckPassword", []byte(password), []byte("hashedpassword")).Return(errors.New("invalid password")).Once()

        _, _, err := usrSrv.Login(email, password)

        assert.Error(t, err)
        assert.Equal(t, "error in server", err.Error())
        pwd.AssertExpectations(t)
    })

    t.Run("Error Query Login", func(t *testing.T) {
        email := "john@example.com"
        password := "password123"

        qry.On("Login", email).Return(users.Users{}, errors.New("login error")).Once()

        _, _, err := usrSrv.Login(email, password)

        assert.Error(t, err)
        assert.Equal(t, "error in server", err.Error())
        qry.AssertExpectations(t)
    })

    t.Run("Error Generating JWT", func(t *testing.T) {
        email := "john@example.com"
        password := "password123"
        userData := users.Users{
            ID:       1,
            Email:    "john@example.com",
            Password: "hashedpassword",
        }

        qry.On("Login", email).Return(userData, nil).Once()
        pwd.On("CheckPassword", []byte(password), []byte("hashedpassword")).Return(nil).Once()
        jwt.On("GenerateJwt", userData.ID).Return("", errors.New("jwt error")).Once()

        _, _, err := usrSrv.Login(email, password)

        assert.Error(t, err)
        assert.Equal(t, "jwt error", err.Error())
        jwt.AssertExpectations(t)
    })
}

func TestGetUser(t *testing.T) {
	qry := mocks.NewQuery(t)  // Mock untuk users.Query
	jwt := mocks.NewJwtUtilityInteface(t) // Mock untuk utils.JwtUtilityInterface
	pwd := mocks.NewPasswordUtilityInterface(t) // Mock untuk utils.PasswordUtilityInterface
	usrSrv := service.NewUserServices(qry, jwt, pwd)

	t.Run("Success Get User", func(t *testing.T) {
		userID := uint(1)
		expectedUser := users.Users{
			ID:    1,
			Username:  "John Doe",
			Email: "john@example.com",
		}

		// Mocking the GetUser query
		qry.On("GetUser", userID).Return(expectedUser, nil).Once()

		result, err := usrSrv.GetUser(userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, result)
		qry.AssertExpectations(t)
	})

	t.Run("Error Get User", func(t *testing.T) {
		userID := uint(1)

		// Mocking the GetUser query to return an error
		qry.On("GetUser", userID).Return(users.Users{}, errors.New("user not found")).Once()

		_, err := usrSrv.GetUser(userID)

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		qry.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
    qry := mocks.NewQuery(t)  // Mock untuk users.Query
    jwt := mocks.NewJwtUtilityInteface(t) // Mock untuk utils.JwtUtilityInterface
    pwd := mocks.NewPasswordUtilityInterface(t) // Mock untuk utils.PasswordUtilityInterface
    usrSrv := service.NewUserServices(qry, jwt, pwd)

    t.Run("Success Update User", func(t *testing.T) {
        userID := uint(1)
        updatedUser := users.Users{
            Username: "Updated Name",
            Password: "newpassword",
        }

        pwd.On("GeneretePassword", "newpassword").Return([]byte("hashednewpassword"), nil).Once()
        qry.On("UpdateUser", userID, mock.Anything).Return(nil).Once()

        err := usrSrv.UpdateUser(userID, updatedUser)

        assert.NoError(t, err)
        pwd.AssertExpectations(t)
        qry.AssertExpectations(t)
    })

    t.Run("Error Hashing Password", func(t *testing.T) {
        userID := uint(1)
        updatedUser := users.Users{
            Username: "Updated Name",
            Password: "newpassword",
        }

        pwd.On("GeneretePassword", "newpassword").Return([]byte{}, errors.New("hash error")).Once()

        err := usrSrv.UpdateUser(userID, updatedUser)

        assert.Error(t, err)
        assert.Equal(t, "hash error", err.Error())
        pwd.AssertExpectations(t)
    })

    t.Run("Error Updating User", func(t *testing.T) {
        userID := uint(1)
        updatedUser := users.Users{
            Username: "Updated Name",
            Password: "newpassword",
        }

        pwd.On("GeneretePassword", "newpassword").Return([]byte("hashednewpassword"), nil).Once()
        qry.On("UpdateUser", userID, mock.Anything).Return(errors.New("update error")).Once()

        err := usrSrv.UpdateUser(userID, updatedUser)

        assert.Error(t, err)
        assert.Equal(t, "error in server", err.Error())
        qry.AssertExpectations(t)
        pwd.AssertExpectations(t)
    })
}


func TestDeleteUser(t *testing.T) {
	qry := mocks.NewQuery(t)  // Mock untuk users.Query
	jwt := mocks.NewJwtUtilityInteface(t) // Mock untuk utils.JwtUtilityInterface
	pwd := mocks.NewPasswordUtilityInterface(t) // Mock untuk utils.PasswordUtilityInterface
	usrSrv := service.NewUserServices(qry, jwt, pwd)

	t.Run("Success Delete User", func(t *testing.T) {
		userID := uint(1)

		// Mocking the DeleteUser query
		qry.On("DeleteUser", userID).Return(nil).Once()

		err := usrSrv.DeleteUser(userID)

		assert.NoError(t, err)
		qry.AssertExpectations(t)
	})

	t.Run("Error Delete User", func(t *testing.T) {
		userID := uint(1)

		// Mocking the DeleteUser query to return an error
		qry.On("DeleteUser", userID).Return(errors.New("delete error")).Once()

		err := usrSrv.DeleteUser(userID)

		assert.Error(t, err)
		assert.Equal(t, "server error", err.Error())
		qry.AssertExpectations(t)
	})
}
