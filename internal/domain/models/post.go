// internal/domain/models/post.go
package models

type Post struct {
	ID       string `gorm:"primaryKey"`
	Title    string
	Content  string
	AuthorID string
}
