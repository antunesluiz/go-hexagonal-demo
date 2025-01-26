// internal/adapters/database/postgres_post_repo.go

package database

import (
	"errors"

	"github.com/antunesluiz/go-hexagonal-demo/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresPostRepo struct {
	DB *gorm.DB
}

func NewPostgresPostRepo(dsn string) (*PostgresPostRepo, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Post{}); err != nil {
		return nil, err
	}

	return &PostgresPostRepo{DB: db}, nil
}

func (r *PostgresPostRepo) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post

	result := r.DB.Find(&posts)

	return posts, result.Error
}

func (r *PostgresPostRepo) GetPostByID(id string) (*models.Post, error) {
	var post models.Post
	result := r.DB.First(&post, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &post, result.Error
}

func (r *PostgresPostRepo) CreatePost(post models.Post) (*models.Post, error) {
	result := r.DB.Create(&post)

	return &post, result.Error
}

func (r *PostgresPostRepo) UpdatePost(post models.Post) (*models.Post, error) {
	result := r.DB.Save(&post)

	return &post, result.Error
}

func (r *PostgresPostRepo) DeletePost(id string) error {
	result := r.DB.Delete(&models.Post{}, "id = ?", id)

	return result.Error
}
