package utils

import "golang.org/x/crypto/bcrypt"

type PasswordUtilityInterface interface {
	GeneretePassword(currentPw string) ([]byte, error)
	CheckPassword(inputPw []byte, currentPw []byte) error
}

type PasswordUtility struct{}

func NewPasswordUtility() PasswordUtilityInterface {
	return &PasswordUtility{}
}

func (pw *PasswordUtility) GeneretePassword(currentPw string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(currentPw), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pw *PasswordUtility) CheckPassword(inputPw []byte, currentPw []byte) error{
	return	bcrypt.CompareHashAndPassword(currentPw, inputPw)
}