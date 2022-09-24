package main

import (
	"encoding/json"
	shortenURL "url-shortener/internal/domain/usecases/url/shorten"
	"url-shortener/internal/infrastructure/adapters/logger"
	"url-shortener/internal/infrastructure/adapters/response"
	shortenURLUseCaseFactory "url-shortener/internal/infrastructure/factories/url/shorten-usecase"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var client = lambda.New(session.New())

func LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log := logger.NewLogger()

	var input shortenURL.Input

	json.Unmarshal([]byte(request.Body), &input)

	log.Info("request to shorten url -> ", input)

	useCase := shortenURLUseCaseFactory.Make(log)

	shortenedURL, appError := useCase.Execute(input)

	if appError != nil {
		log.Info("", appError)
		return response.BadRequest(response.BadRequestInput{Message: appError.Message, StatusCode: int(appError.StatusCode)})
	}

	return response.Ok(shortenedURL)
}

func main() {
	runtime.Start(LambdaHandler)
}
