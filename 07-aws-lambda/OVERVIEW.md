# 07 - AWS Lambda

This Go program implements an AWS Lambda function that processes events containing name and age, and returns a formatted message. It demonstrates serverless computing with AWS Lambda.

### 🔧 Imports

```go
import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)
```

- `fmt`: For string formatting.
- `github.com/aws/aws-lambda-go/lambda`: AWS Lambda Go SDK.

### 🎬 Data Structures

```go
type MyEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MyResponse struct {
	Message string `json:"message"`
}
```

- `MyEvent`: Input event structure with Name and Age.
- `MyResponse`: Output response with Message.

### 🧩 Handler Function

```go
func handleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{
		Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age),
	}, nil
}
```

- Takes a MyEvent, formats a message, returns MyResponse.
- No error handling needed for this simple example.

### 🚀 Main Function

```go
func main() {
	lambda.Start(handleLambdaEvent)
}
```

- Starts the Lambda runtime with the handler function.

### 📝 Usage

- Deploy to AWS Lambda.
- Invoke with JSON: `{"name": "John", "age": 30}`
- Returns: `{"message": "John is 30 years old!"}`

### 📦 Deployment

This project includes:
- `template.yaml`: SAM template for deployment
- `event.json`: Sample event for testing
- `trust-policy.json`: IAM trust policy