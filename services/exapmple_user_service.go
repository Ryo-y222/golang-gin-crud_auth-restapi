package services

import (
	"gin-fleamarket/repositories"
)

type UserService interface {
	GetGreetingUserNames() []string
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	// fmt.Printf("%#v\n", repo)
	// spew.Dump(repo)
	return &userService{repo}
}

func (s *userService) GetGreetingUserNames() []string {
	names := s.repo.GetUserNames()
	var result []string
	for _, name := range names {
		result = append(result, "Hello,"+name)
	}
	return result
}
