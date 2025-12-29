# 07 - AWS Lambda

An AWS Lambda function written in Go that processes events with name and age data and returns a formatted message. This project demonstrates serverless computing with AWS Lambda.

## Features

- Processes JSON events with name and age
- Returns formatted response message
- Uses AWS Lambda Go SDK

## Project Structure

- `main.go`: Lambda handler function
- `template.yaml`: AWS SAM deployment template
- `event.json`: Sample input event for testing
- `trust-policy.json`: IAM role trust policy
- `go.mod`, `go.sum`: Go module files

## How to Deploy

1. Navigate to the project directory:
   ```sh
   cd "07-aws-lambda"
   ```
2. Build the function:
   ```sh
   GOOS=linux GOARCH=amd64 go build -o main main.go
   ```
3. Zip the binary:
   ```sh
   zip lambda-function.zip main
   ```
4. Deploy using AWS SAM or directly to Lambda

## Testing Locally

Use the AWS SAM CLI to test locally:
```sh
sam local invoke -e event.json
```

## API Gateway Integration

Can be triggered via API Gateway with JSON payload:
```json
{
  "name": "Alice",
  "age": 25
}
```

Response:
```json
{
  "message": "Alice is 25 years old!"
}
```

---

This project is part of my Go learning journey!