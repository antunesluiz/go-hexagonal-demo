// internal/adapters/database/postgres_user_repo.go
package database

import (
	"errors"

	"github.com/antunesluiz/go-hexagonal-demo/internal/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresUserRepo struct {
	DB *gorm.DB
}

func NewPostgresUserRepo(dsn string) (*PostgresUserRepo, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return &PostgresUserRepo{DB: db}, nil
}

func (r *PostgresUserRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User

	result := r.DB.Find(&users)

	return users, result.Error
}

func (r *PostgresUserRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, result.Error
}

func (r *PostgresUserRepo) CreateUser(user models.User) (*models.User, error) {
	result := r.DB.Create(&user)

	return &user, result.Error
}

func (r *PostgresUserRepo) UpdateUser(user models.User) (*models.User, error) {
	result := r.DB.Save(&user)

	return &user, result.Error
}

func (r *PostgresUserRepo) DeleteUser(id string) error {
	result := r.DB.Delete(&models.User{}, "id = ?", id)

	return result.Error
}
