package main

import (
	"fmt"

	"github.com/stephane-nguyen/Bubble-tea/pkg/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

func main() {
	region := "eu-west-3"
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		fmt.Printf("Error creating AWS session: %v\n", err)
		return
	}

	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

const tableName = "BubbleTea"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetBoba(req, tableName, dynaClient)
	case "POST":
		return handlers.CreateBoba(req, tableName, dynaClient)
	case "PUT":
		return handlers.UpdateBoba(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeleteBoba(req, tableName, dynaClient)
	default:
		return handlers.UnhandledMethod()
	}
}
