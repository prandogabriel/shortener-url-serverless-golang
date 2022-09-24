package dynamo_shortened_url_repository

import (
	"errors"
	"url-shortener/internal/domain/entities"
	"url-shortener/internal/domain/ports/repositories"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dynamoDB struct {
	Api    dynamodbiface.DynamoDBAPI
	Table  string
	HashId string
}

func NewShortenedUrlRepository() repositories.ShortenedUrlRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	return &dynamoDB{
		Api:    svc,
		Table:  "ShortenedUrl",
		HashId: "id",
	}
}

func (repo *dynamoDB) Save(entity *entities.ShortenedUrl) error {

	item, err := dynamodbattribute.MarshalMap(entity)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: &repo.Table,
		Item:      item,
	}

	_, err = repo.Api.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}

func (repo *dynamoDB) FindById(id *string) (*entities.ShortenedUrl, error) {

	getInput := &dynamodb.GetItemInput{
		TableName: aws.String(repo.Table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(*id),
			},
		},
	}

	out, err := repo.Api.GetItem(getInput)

	if err != nil {
		return &entities.ShortenedUrl{}, err
	}

	item := entities.ShortenedUrl{}

	err = dynamodbattribute.UnmarshalMap(out.Item, &item)

	if err != nil {
		return &entities.ShortenedUrl{}, errors.New("Error on get shortenedUrls")
	}

	return &item, nil
}
