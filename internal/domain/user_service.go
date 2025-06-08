package domain

type UserService interface {
	GetAllUsers() []User
	CreateUser(user User) User
}
