package boba

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	ErrorFailedToFetchRecord      = "failed to fetch record"
	ErrorFailtedToUnmarshalRecord = "failed to unmarshal record"
	ErrorInvalidBobaData          = "invalid boba data"
	ErrorCouldNotMarshalItem      = "could not marshal item"
	ErrorCouldNotDeleteItem       = "could not delete item"
	ErrorCouldNotPutItem          = "could not put item"
	ErrorBobaAlreadyExists        = "boba already exists"
	ErrorBobaDoesNotExist         = "boba does not exist"
	ErrorInvalidBobaName          = "invalid boba name"
)

type Boba struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Flavor string `json:"flavor"`
}

func FetchBoba(name, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*Boba, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
	}

	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	item := new(Boba)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailtedToUnmarshalRecord)
	}

	return item, nil
}

func FetchBobas(tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*[]Boba, error) {
	items := new([]Boba)
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, items)
	if err != nil {
		return nil, errors.New(ErrorFailtedToUnmarshalRecord)
	}

	return items, nil
}

func CreateBoba(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*Boba, error) {
	var boba Boba

	if err := json.Unmarshal([]byte(req.Body), &boba); err != nil {
		return nil, errors.New(ErrorInvalidBobaData)
	}

	currentBoba, _ := FetchBoba(boba.Name, tableName, dynaClient)
	if currentBoba != nil {
		return nil, errors.New(ErrorBobaAlreadyExists)
	}

	av, err := dynamodbattribute.MarshalMap(boba)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotPutItem)
	}

	return &boba, nil
}

func UpdateBoba(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*Boba, error) {
	var boba Boba

	if err := json.Unmarshal([]byte(req.Body), &boba); err != nil {
		return nil, errors.New(ErrorInvalidBobaName)
	}

	currentBoba, _ := FetchBoba(boba.Name, tableName, dynaClient)
	if currentBoba != nil {
		return nil, errors.New(ErrorBobaDoesNotExist)
	}

	av, err := dynamodbattribute.MarshalMap(boba)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotPutItem)
	}

	return &boba, nil
}

func DeleteBoba(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) error {
	name := req.QueryStringParameters["name"]
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(name),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := dynaClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
