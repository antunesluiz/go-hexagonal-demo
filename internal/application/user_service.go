// internal/application/user_service.go
package application

import (
	"github.com/antunesluiz/go-hexagonal-demo/internal/domain/models"
	"github.com/antunesluiz/go-hexagonal-demo/internal/domain/ports"
	"github.com/google/uuid"
)

type UserService struct {
	Repo ports.UserService
}

func NewUserService(repo ports.UserService) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) CreateUser(name, email string) (*models.User, error) {
	user := models.User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(id, name, email string) (*models.User, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	return s.Repo.UpdateUser(*user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.Repo.DeleteUser(id)
}
