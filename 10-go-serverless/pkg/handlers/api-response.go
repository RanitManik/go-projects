package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func APIResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "failed to marshal response body",
		}, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonBody),
	}, nil
}
