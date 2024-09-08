package users

type Users struct {
	ID       uint
	Username string
	Email    string
	Password string
}

type Handler interface {
}

type Query interface {
	Register(newUser Users) error
	Login(email string) (Users, error)
	UpdateUser(id uint, updateUser Users) error
	DeleteUser(id uint) error
}

type Service interface {
	Register(newUser Users) error
	Login(email string, password string) (Users, string, error)
	UpdateUser(id uint, updateUser Users) error
	DeleteUser(id uint) error
}