package shorten_url

import (
	"url-shortener/internal/domain/ports/logger"
	shortenUrl "url-shortener/internal/domain/usecases/shorten-url"
	shortenedRepo "url-shortener/internal/infrastructure/adapters/repositories/shortened-url"
)

func Make(log logger.Logger) shortenUrl.ShortenURLUseCase {
	dynamoShortenedUrlRepository := shortenedRepo.NewShortenedUrlRepository()

	return shortenUrl.New(dynamoShortenedUrlRepository, log)
}
