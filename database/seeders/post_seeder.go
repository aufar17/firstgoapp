package seeders
	
import (
  "github.com/goravel/framework/facades"
  "goravel/app/models"

)
type PostSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *PostSeeder) Signature() string {
	return "PostSeeder"
}

// Run executes the seeder logic.
func (s *PostSeeder) Run() error {
	return facades.Orm().Query().Create(&models.Post{
		Title: "First App with GO Lang and Goravel",
		Body: "Goravel test first post ",
		Author: "Super Pakcoy",
	})
}
