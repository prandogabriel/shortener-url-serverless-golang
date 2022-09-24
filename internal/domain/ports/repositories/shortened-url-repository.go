package repositories

import "url-shortener/internal/domain/entities"

type ShortenedUrlRepository interface {
	Save(entity *entities.ShortenedUrl) error
	FindById(id *string) (*entities.ShortenedUrl, error)
}
