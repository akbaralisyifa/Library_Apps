package service

import (
	"errors"
	"library/internal/features/users"
	"library/internal/utils"
	"log"
)

type UserServices struct {
	qry users.Query
	jwt utils.JwtUtilityInteface
	pwd utils.PasswordUtilityInterface
}

func NewUserServices(q users.Query, j utils.JwtUtilityInteface, p utils.PasswordUtilityInterface) users.Service {
	return &UserServices{
		qry: q,
		jwt: j,
		pwd: p,
	}
}


func (us *UserServices) Register(newUser users.Users) error {

	// hashing password
	hashPw, err := us.pwd.GeneretePassword(newUser.Password);
	if err != nil {
		log.Print("Register Generete Password Error", err.Error())
		return err
	}

	newUser.Password = string(hashPw)

	err = us.qry.Register(newUser);
	if err != nil {
		log.Print("Register sql Error :", err.Error())
		return errors.New("error in ervice")
	}

	return	nil
}

func (us *UserServices) Login(email string, password string) (users.Users, string, error) {
	result, err := us.qry.Login(email);
	if err != nil {
		log.Print("Login sql Error :", err.Error())
		return users.Users{}, "", errors.New("error in server")
	}

	// cek password
	err = us.pwd.CheckPassword([]byte(password), []byte(result.Password))
	if err != nil {
		log.Print("check password error", err.Error())
		return users.Users{}, "", errors.New("error in server")
	}

	// Generete Token
	token, err := us.jwt.GenerateJwt(result.ID);
	if err != nil {
		log.Print("generete token error", err.Error())
		return users.Users{}, "", errors.New("error in server")
	}

	return result, token, nil
}


func (us *UserServices) UpdateUser(id uint, updateUser users.Users) error {

	// check pw update
	if updateUser.Password != "" {
		hashPw, err := us.pwd.GeneretePassword(updateUser.Password)
		if err != nil {
			log.Print("update generete password error", err.Error())
			return errors.New("update generete password error")
		}
		updateUser.Password = string(hashPw)
	}

	err := us.qry.UpdateUser(id, updateUser)
	if err != nil {
		log.Print("update user error", err.Error())
		return errors.New("server error")
	}

	return nil;
}

func (us *UserServices) DeleteUser(id uint) error {
	err := us.qry.DeleteUser(id);

	if err != nil {
		log.Print("delete user error", err.Error())
		return errors.New("server error")
	}

	return nil;
}