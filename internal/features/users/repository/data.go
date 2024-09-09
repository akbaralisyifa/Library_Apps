package repository

import (
	rrepo "library/internal/features/recomendation/repository"
	"library/internal/features/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Role     string
	Recomendation []rrepo.Recomendation `gorm:"foreignKey:UserID"`
}

func ToUserQuery(input users.Users) Users{
	return Users{
		Username: input.Username,
		Email: 	  input.Email,
		Password: input.Password,
		Role:     input.Role,
	}
}

func (us *Users) ToUserEntity() users.Users{
	return users.Users{
		ID:		  us.ID,	
		Username: us.Username,
		Email: 	  us.Email,
		Password: us.Password,
		Role:     us.Role,
	}
}