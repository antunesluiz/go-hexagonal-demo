// internal/application/post_service.go

package application

import (
	"github.com/antunesluiz/go-hexagonal-demo/internal/domain/models"
	"github.com/antunesluiz/go-hexagonal-demo/internal/domain/ports"
	"github.com/google/uuid"
)

type PostService struct {
	Repo ports.PostService
}

func NewPostService(repo ports.PostService) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) GetAllPosts() ([]models.Post, error) {
	return s.Repo.GetAllPosts()
}

func (s *PostService) GetPostByID(id string) (*models.Post, error) {
	return s.Repo.GetPostByID(id)
}

func (s *PostService) CreatePost(title, content, authorID string) (*models.Post, error) {
	post := models.Post{
		ID:       uuid.New().String(),
		Title:    title,
		Content:  content,
		AuthorID: authorID,
	}

	return s.Repo.CreatePost(post)
}

func (s *PostService) UpdatePost(id, title, content, authorID string) (*models.Post, error) {
	post, err := s.Repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	if title != "" {
		post.Title = title
	}
	if content != "" {
		post.Content = content
	}
	if authorID != "" {
		post.AuthorID = authorID
	}

	return s.Repo.UpdatePost(*post)
}

func (s *PostService) DeletePost(id string) error {
	return s.Repo.DeletePost(id)
}
