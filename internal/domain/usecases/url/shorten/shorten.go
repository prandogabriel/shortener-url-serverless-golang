package shorten_url

import (
	"url-shortener/internal/domain/entities"
	"url-shortener/internal/domain/errors"
	"url-shortener/internal/domain/ports/logger"
	"url-shortener/internal/domain/ports/repositories"
	"url-shortener/internal/domain/utils/date"
	"url-shortener/internal/domain/utils/uid"
)

type ShortenURLUseCase interface {
	Execute(input Input) (entities.ShortenedUrl, *errors.AppError)
}

type useCase struct {
	shortenedUrlRepository repositories.ShortenedUrlRepository
	log                    logger.Logger
}

func New(shortenedUrlRepository repositories.ShortenedUrlRepository, log logger.Logger) ShortenURLUseCase {
	return &useCase{shortenedUrlRepository: shortenedUrlRepository, log: log}
}

func (uc *useCase) Execute(input Input) (entities.ShortenedUrl, *errors.AppError) {

	shortenedUrl := buildEntity(input)

	uc.log.Info("Creating a new shortened URL ", *shortenedUrl)

	err := uc.shortenedUrlRepository.Save(shortenedUrl)

	if err != nil {
		return entities.ShortenedUrl{}, errors.InternalError(err)
	}

	return *shortenedUrl, nil
}

func buildEntity(input Input) *entities.ShortenedUrl {
	id := uid.New()
	now := date.Now().DynamoFormat()

	return &entities.ShortenedUrl{
		ID:              id,
		Name:            input.Name,
		OriginalURL:     input.URL,
		RecoveriesCount: 0,
		CreateBy:        input.CreateBy,
		CreateDate:      now,
		UpdateDate:      now,
	}
}
