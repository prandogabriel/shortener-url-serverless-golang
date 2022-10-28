package shortened_url_repository

import (
	"errors"
	"url-shortener/internal/domain/entities"
	"url-shortener/internal/domain/ports/repositories"
)

type mockDB struct {
	items []entities.ShortenedUrl
}

func NewShortenedUrlRepositoryMock(items []entities.ShortenedUrl) repositories.ShortenedUrlRepository {

	return &mockDB{
		items: items,
	}
}

func (repo *mockDB) Save(entity *entities.ShortenedUrl) error {
	repo.items = append(repo.items, *entity)

	return nil
}

func (repo *mockDB) FindById(id *string) (*entities.ShortenedUrl, error) {
	item := entities.ShortenedUrl{}

	for _, v := range repo.items {
		if v.ID == *id {
			item = v
			break
		}
	}

	if (item == entities.ShortenedUrl{}) {
		return &item, errors.New("Not found")
	}

	return &item, nil
}
