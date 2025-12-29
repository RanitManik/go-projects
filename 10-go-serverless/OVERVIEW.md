# 10 - Go Serverless

This Go program implements a serverless REST API using AWS Lambda and DynamoDB for managing user records. It handles CRUD operations through API Gateway integration.

### 🔧 Imports

```go
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
```

- `context`, `log`, `os`: Standard libraries.
- `github.com/aws/aws-lambda-go/*`: AWS Lambda SDK.
- `github.com/aws/aws-sdk-go-v2/*`: AWS SDK v2 for DynamoDB.

### 🎬 Data Structures

```go
type User struct {
	Email     string `json:"email" dynamodbav:"email"`
	FirstName string `json:"firstName" dynamodbav:"firstName"`
	LastName  string `json:"lastName" dynamodbav:"lastName"`
}
```

- `User` represents a user with Email (primary key), FirstName, LastName.

### 📦 Database Connection

```go
var dynaClient *dynamodb.Client

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
```

- Initializes DynamoDB client with AWS region from environment.

### 🧩 Handler Function

```go
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
```

- Routes HTTP methods to appropriate handlers.

### 📋 Handlers

#### ✅ Get Users

```go
func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		// ... return single user ...
	}

	result, err := user.FetchUsers(tableName, dynaClient)
	// ... return all users ...
}
```

- GET with email param fetches single user, without param fetches all.

#### ➕ Create User

```go
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(req, tableName, dynaClient)
	if err != nil {
		return APIResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return APIResponse(http.StatusCreated, result)
}
```

- Creates new user from JSON body.

#### ✏️ Update User

```go
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	result, err := user.UpdateUser(req, tableName, dynaClient)
	// ... error handling ...
	return APIResponse(http.StatusOK, result)
}
```

- Updates existing user.

#### ❌ Delete User

```go
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
	err := user.DeleteUser(req, tableName, dynaClient)
	if err != nil {
		return APIResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return APIResponse(http.StatusOK, nil)
}
```

- Deletes user by email.

### 🚀 Deployment

- Deploy to AWS Lambda with API Gateway trigger.
- Requires DynamoDB table `LambdaInGoUser` with email as primary key.
- Set `AWS_REGION` environment variable.

### 📝 Summary of API Endpoints

| Method | Endpoint      | Description                 |
|--------|---------------|-----------------------------|
| GET    | `/`           | List all users              |
| GET    | `/?email=...` | Get a user by email         |
| POST   | `/`           | Create a new user           |
| PATCH  | `/`           | Update an existing user     |
| DELETE | `/`           | Delete a user by email      |