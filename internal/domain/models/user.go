// internal/domain/models/user.go
package models

type User struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"uniqueIndex"`
}
