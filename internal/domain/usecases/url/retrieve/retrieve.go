package retrieve_url

import (
	"url-shortener/internal/domain/entities"
	"url-shortener/internal/domain/errors"
	"url-shortener/internal/domain/ports/logger"
	"url-shortener/internal/domain/ports/repositories"
	"url-shortener/internal/domain/utils/date"
)

type RetrieveURLUseCase interface {
	Execute(input Input) (*entities.ShortenedUrl, *errors.AppError)
}

type useCase struct {
	shortenedUrlRepository repositories.ShortenedUrlRepository
	log                    logger.Logger
}

func New(shortenedUrlRepository repositories.ShortenedUrlRepository, log logger.Logger) RetrieveURLUseCase {
	return &useCase{shortenedUrlRepository: shortenedUrlRepository, log: log}
}

func (uc *useCase) Execute(input Input) (*entities.ShortenedUrl, *errors.AppError) {
	shortenedUrl, err := uc.shortenedUrlRepository.FindById(&input.ShortenedId)
	if err != nil {
		uc.log.Error("ShortenedUrl not found", err)
		return &entities.ShortenedUrl{}, errors.BadRequest("ShortenedUrl not found", err)
	}

	increaseRetrieveCount(shortenedUrl, uc.shortenedUrlRepository)

	return shortenedUrl, nil
}

func increaseRetrieveCount(shortenedUrl *entities.ShortenedUrl, repository repositories.ShortenedUrlRepository) {
	(*shortenedUrl).RecoveriesCount++
	(*shortenedUrl).UpdateDate = date.Now().DynamoFormat()

	repository.Save(shortenedUrl)
}
