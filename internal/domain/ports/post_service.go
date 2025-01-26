// internal/domain/ports/post_service.go
package ports

import "github.com/antunesluiz/go-hexagonal-demo/internal/domain/models"

type PostService interface {
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id string) (*models.Post, error)
	CreatePost(post models.Post) (*models.Post, error)
	UpdatePost(post models.Post) (*models.Post, error)
	DeletePost(id string) error
}
