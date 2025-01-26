// internal/domain/ports/user_service.go
package ports

import "github.com/antunesluiz/go-hexagonal-demo/internal/domain/models"

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
	DeleteUser(id string) error
}
