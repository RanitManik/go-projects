# 10 - Go Serverless

A serverless REST API built with AWS Lambda and DynamoDB for managing user records. This project demonstrates serverless architecture with full CRUD operations.

## Features

- List all users
- Get a single user by email
- Create a new user
- Update an existing user
- Delete a user
- Email validation
- DynamoDB integration

## Project Structure

- `cmd/main.go`: Main Lambda handler and routing
- `pkg/handlers/handlers.go`: HTTP method handlers
- `pkg/user/user.go`: User model and DynamoDB operations
- `pkg/validators/email-validator.go`: Email validation logic
- `go.mod`, `go.sum`: Go module files

## Prerequisites

- AWS account with Lambda and DynamoDB access
- DynamoDB table: `LambdaInGoUser`
  - Primary key: `email` (String)
- API Gateway for HTTP triggers
- Environment variable: `AWS_REGION`

## Deployment

1. Build for Linux:
   ```sh
   GOOS=linux GOARCH=amd64 go build -o main cmd/main.go
   ```
2. Zip the binary:
   ```sh
   zip lambda-function.zip main
   ```
3. Deploy to AWS Lambda via AWS Console, CLI, or SAM

## API Endpoints

- `GET    /` - List all users
- `GET    /?email={email}` - Get a user by email
- `POST   /` - Create a new user (JSON body)
- `PATCH  /` - Update a user (JSON body with email)
- `DELETE /` - Delete a user (JSON body with email)

## Example User JSON

```json
{
  "email": "john@example.com",
  "firstName": "John",
  "lastName": "Doe"
}
```

## DynamoDB Table Schema

- Table Name: `LambdaInGoUser`
- Key: `email` (String)

## Environment Variables

- `AWS_REGION`: AWS region (e.g., us-east-1)

---

This project is part of my Go learning journey!