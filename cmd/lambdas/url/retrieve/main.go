package main

import (
	retrieveURL "url-shortener/internal/domain/usecases/url/retrieve"
	"url-shortener/internal/infrastructure/adapters/logger"
	"url-shortener/internal/infrastructure/adapters/response"
	retrieveURLUseCaseFactory "url-shortener/internal/infrastructure/factories/url/retrieve-usecase"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var client = lambda.New(session.New())

func LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log := logger.NewLogger()

	var input retrieveURL.Input

	input.ShortenedId = request.PathParameters["shortenedId"]

	log.Info("request to retrieve url -> ", input)

	useCase := retrieveURLUseCaseFactory.Make(log)

	shortenedURL, appError := useCase.Execute(input)

	if appError != nil {
		log.Info("", appError)
		return response.BadRequest(response.BadRequestInput{Message: appError.Message, StatusCode: int(appError.StatusCode)})
	}

	return response.Redirect(shortenedURL.OriginalURL)
}

func main() {
	runtime.Start(LambdaHandler)
}
