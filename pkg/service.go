package pkg

import (
	"errors"
)

type Service struct {
	users []Database
}

func NewService() *Service {
	s := &Service{
		users: make([]Database, 0),
	}
	s.DummyUsers()
	return s
}

func (s *Service) DummyUsers() {
	s.users = append(s.users, Database{Name: "Turab", Email: "turab@gmail.com", Gender: "male"})
	s.users = append(s.users, Database{Name: "Saif", Email: "saif@gmail.com", Gender: "male"})
}

func (s *Service) CreateUserdata(name, email, gender string) (Database, error) {
	if name == "" || email == "" {
		return Database{}, errors.New("name and email cannot be empty")
	}
	user := Database{
		Name:   name,
		Email:  email,
		Gender: gender,
	}
	s.users = append(s.users, user)
	return user, nil
}
func (s *Service) GetAllUsers() []Database {
	return s.users
}
