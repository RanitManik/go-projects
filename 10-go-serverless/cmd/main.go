package main

import (
	"context"
	"log"
	"os"

	"github.com/RanitManik/go-projects/10-go-serverless/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var dynaClient *dynamodb.Client

const tableName = "LambdaInGoUser"

func main() {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		log.Fatal("AWS_REGION is not set")
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatal(err)
	}

	dynaClient = dynamodb.NewFromConfig(cfg)

	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser(req, tableName, dynaClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynaClient)
	case "PATCH":
		return handlers.UpdateUser(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynaClient)
	default:
		return handlers.DefaultHandler()
	}
}
