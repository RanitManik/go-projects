package user

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/RanitManik/go-projects/10-go-serverless/pkg/validators"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var (
	ErrorFailedToFetchRecord     = errors.New("failed to fetch record")
	ErrorFailedToUnmarshalRecord = errors.New("failed to unmarshal record")
	ErrorUserNotFound            = errors.New("user not found")
	ErrorInvalidUserData         = errors.New("invalid user data")
	ErrorInvalidEmail            = errors.New("invalid email")
	ErrorFailedToDeleteItem      = errors.New("failed to delete item")
	ErrorFailedToDynamoPutItem   = errors.New("failed to dynamo put item")
	ErrorUserAlreadyExists       = errors.New("user already exists")
	ErrorUserDoesNotExist        = errors.New("user does not exist")
	ErrorInternalServerError     = errors.New("internal server error")
)

type User struct {
	Email     string `json:"email" dynamodbav:"email"`
	FirstName string `json:"firstName" dynamodbav:"firstName"`
	LastName  string `json:"lastName" dynamodbav:"lastName"`
}

func FetchUsers(tableName string, dynaClient *dynamodb.Client) ([]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.Scan(context.TODO(), input)
	if err != nil {
		return nil, ErrorFailedToFetchRecord
	}

	var users []User
	if err := attributevalue.UnmarshalListOfMaps(result.Items, &users); err != nil {
		return nil, ErrorFailedToUnmarshalRecord
	}

	return users, nil
}

func FetchUser(email string, tableName string, dynamoClient *dynamodb.Client) (*User, error) {
	if !validators.IsEmailValid(email) {
		return nil, ErrorInvalidEmail
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{
				Value: email,
			},
		},
	}

	result, err := dynamoClient.GetItem(context.TODO(), input)
	if err != nil {
		return nil, ErrorFailedToFetchRecord
	}

	if result.Item == nil {
		return nil, ErrorUserNotFound
	}

	var user User
	if err := attributevalue.UnmarshalMap(result.Item, &user); err != nil {
		return nil, ErrorFailedToUnmarshalRecord
	}

	return &user, nil
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient *dynamodb.Client) (*User, error) {
	var user User

	if err := json.Unmarshal([]byte(req.Body), &user); err != nil {
		return nil, ErrorFailedToUnmarshalRecord
	}

	if !validators.IsEmailValid(user.Email) {
		return nil, ErrorInvalidEmail
	}

	fetchedUser, err := FetchUser(user.Email, tableName, dynamoClient)
	if err != nil && err != ErrorUserNotFound {
		return nil, ErrorInternalServerError
	}
	if fetchedUser != nil {
		return nil, ErrorUserAlreadyExists
	}

	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		return nil, ErrorFailedToUnmarshalRecord
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynamoClient.PutItem(context.TODO(), input)
	if err != nil {
		return nil, ErrorFailedToDynamoPutItem
	}

	return &user, nil
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient *dynamodb.Client) (*User, error) {
	var user User

	if err := json.Unmarshal([]byte(req.Body), &user); err != nil {
		return nil, ErrorFailedToUnmarshalRecord
	}

	if !validators.IsEmailValid(user.Email) {
		return nil, ErrorInvalidEmail
	}

	_, err := FetchUser(user.Email, tableName, dynamoClient)
	if err == ErrorUserNotFound {
		return nil, ErrorUserDoesNotExist
	}
	if err != nil {
		return nil, ErrorInternalServerError
	}

	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		return nil, ErrorFailedToUnmarshalRecord
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynamoClient.PutItem(context.TODO(), input)
	if err != nil {
		return nil, ErrorFailedToDynamoPutItem
	}

	return &user, nil
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynamoClient *dynamodb.Client) error {
	email := req.QueryStringParameters["email"]
	if !validators.IsEmailValid(email) {
		return ErrorInvalidEmail
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{
				Value: email,
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := dynamoClient.DeleteItem(context.TODO(), input)
	if err != nil {
		return ErrorFailedToDeleteItem
	}
	return nil
}
