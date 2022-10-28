package retrieve_url

import (
	"testing"
	"url-shortener/internal/domain/entities"
	"url-shortener/internal/infrastructure/adapters/logger"
	shortenedUrlRepository "url-shortener/internal/infrastructure/adapters/repositories/shortened-url"
	"url-shortener/pkg/utils/date"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("Successful use case find url", func(t *testing.T) {
		items := make([]entities.ShortenedUrl, 10)

		now := date.Now().DynamoFormat()
		item := entities.ShortenedUrl{ID: "12asfasr", Name: "test", OriginalURL: "https://google.com", RecoveriesCount: 1, CreateBy: "gabriel prando", CreateDate: now, UpdateDate: now}
		items = append(items, item)

		useCase := New(shortenedUrlRepository.NewShortenedUrlRepositoryMock(items), logger.NewLogger())

		input := Input{ShortenedId: item.ID}

		entity, err := useCase.Execute(input)

		assert.Nil(t, err)
		assert.Equal(t, item.CreateDate, entity.CreateDate)
	})

	t.Run("Unsuccessful use case find url", func(t *testing.T) {
		items := make([]entities.ShortenedUrl, 10)

		useCase := New(shortenedUrlRepository.NewShortenedUrlRepositoryMock(items), logger.NewLogger())

		input := Input{ShortenedId: ""}

		_, err := useCase.Execute(input)

		assert.NotNil(t, err)
	})

}
