package repository

import (
	brepo "library/internal/features/borrowed/repository"
	rrepo "library/internal/features/recomendation/repository"
	"library/internal/features/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Recomendation []rrepo.Recomendation `gorm:"foreignKey:UserID"`
	Borrowed	  []brepo.Borrowed `gorm:"foreignKey:UserID"`
}


func ToUserQuery(input users.Users) Users{
	return Users{
		Username: input.Username,
		Email: 	  input.Email,
		Password: input.Password,
	}
}

func (us *Users) ToUserEntity() users.Users{
	return users.Users{
		ID:		  us.ID,	
		Username: us.Username,
		Email: 	  us.Email,
		Password: us.Password,
	}
}