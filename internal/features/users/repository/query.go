package repository

import (
	"library/internal/features/users"

	"gorm.io/gorm"
)

type UserModels struct {
	db *gorm.DB
}

func NewUserModels(connect *gorm.DB) users.Query {
	return	&UserModels{
		db: connect,
	}
}


func (um *UserModels) Register(newUser users.Users) error {
	cnvData := ToUserQuery(newUser)
	err := um.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	return nil;
}

func (um *UserModels) Login(email string) (users.Users, error){
	var result Users;
	err := um.db.Where("email = ?", email).First(&result).Error

	if err != nil {
		return users.Users{}, err
	}

	return result.ToUserEntity(), nil
}


func (um *UserModels) UpdateUser(id uint, updateUser users.Users) error {
	cnvData := ToUserQuery(updateUser);
	
	qry := um.db.Where("id = ?", id).Updates(&cnvData);

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}


func (um *UserModels) DeleteUser(id uint) error {
	qry := um.db.Where("id = ?", id).Delete(&Users{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}