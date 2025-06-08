package domain

type UserRepository interface {
	FindAll() []User
	Save(user User) User
}
