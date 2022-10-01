package retrieve_url

import (
	"url-shortener/internal/domain/ports/logger"
	retrieveUrl "url-shortener/internal/domain/usecases/retrieve-url"
	shortenedRepo "url-shortener/internal/infrastructure/adapters/repositories/shortened-url"
)

func Make(log logger.Logger) retrieveUrl.RetrieveURLUseCase {
	dynamoShortenedUrlRepository := shortenedRepo.NewShortenedUrlRepository()

	return retrieveUrl.New(dynamoShortenedUrlRepository, log)
}
