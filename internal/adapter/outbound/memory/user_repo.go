package memory

import "github.com/santigorbe/hexagonal_arq/internal/domain"

type InMemoryUserRepo struct {
	users  []domain.User
	nextID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users:  []domain.User{},
		nextID: 1,
	}
}

func (r *InMemoryUserRepo) FindAll() []domain.User {
	return r.users
}

func (r *InMemoryUserRepo) Save(user domain.User) domain.User {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return user
}
