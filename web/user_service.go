package main

type UserInterface interface {
	GetUserByEmail(email string) (*User, error)
	InsertUser(name string, email string, hashedPassword string) error
}

type UserService struct {
	userRepo UserInterface
}

func NewUserService(userRepo UserInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(name string, email string, hashedPassword string) error {
	return s.userRepo.InsertUser(name, email, hashedPassword)
}
