package seeders

import (
	"github.com/go-faker/faker/v4"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type PostSeeder struct{}

func (s *PostSeeder) Signature() string {
	return "PostSeeder"
}

func (s *PostSeeder) Run() error {
	posts := make([]models.Post, 0, 1000)

	for i := 0; i < 10000; i++ {
		posts = append(posts, models.Post{
			Title:  faker.Sentence(),
			Body:   faker.Paragraph(),
			Author: faker.Name(),
		})
	}

	// Insert batch
	return facades.Orm().Query().Create(&posts)
}
