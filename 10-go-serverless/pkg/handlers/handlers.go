package handlers

import (
	"net/http"

	"github.com/RanitManik/go-projects/10-go-serverless/pkg/user"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
    ErrorMsg *string `json:"error,omitempty"` 
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return APIResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return APIResponse(http.StatusOK, result)
	}

	result, err := user.FetchUsers(tableName, dynaClient)
	if err != nil {
		return APIResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return APIResponse(http.StatusOK, result)

}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(req, tableName, dynaClient)
	if err != nil {
		return APIResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return APIResponse(http.StatusCreated, result)
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	result, err := user.UpdateUser(req, tableName, dynaClient)
	if err != nil {
		return APIResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return APIResponse(http.StatusOK, result)
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	err := user.DeleteUser(req, tableName, dynaClient)
	if err != nil {
		return APIResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return APIResponse(http.StatusOK, nil)
}

func DefaultHandler() (*events.APIGatewayProxyResponse, error) {
	return APIResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
