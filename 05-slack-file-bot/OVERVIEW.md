# 05 - Slack File Bot

This Go program implements a Slack bot that uploads a markdown file to a specified Slack channel. It uses the Slack Go SDK for file uploads.

### 🔧 Imports

```go
import (
	"bytes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)
```

- `bytes`, `fmt`, `os`: Standard libraries for file handling and output.
- `github.com/joho/godotenv`: For loading environment variables.
- `github.com/slack-go/slack`: Official Slack Go SDK.

### 🧩 Main Functionality

#### File Upload

```go
data, err := os.ReadFile("SampleFile.md")
if err != nil {
	panic(err)
}
if len(data) == 0 {
	panic("File is empty")
}

api := slack.New(token)

params := slack.UploadFileV2Parameters{
	Channel:  channel,
	Filename: "sample-file.md",
	Title:    "Same Markdown File Upload",
	Reader:   bytes.NewReader(data),
	FileSize: len(data),
}

file, err := api.UploadFileV2(params)
if err != nil {
	panic(err)
}

fmt.Printf("Uploaded\nID: %s\nTitle: %s\n", file.ID, file.Title)
```

- Reads the `SampleFile.md` file.
- Creates a Slack API client.
- Uploads the file to the specified channel using UploadFileV2.
- Prints the uploaded file details.

### 🚀 Main Function

```go
func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env")
	}

	token := os.Getenv("SLACK_BOT_TOKEN")
	channel := os.Getenv("CHANNEL_ID")

	if token == "" || channel == "" {
		panic("Missing SLACK_BOT_TOKEN or CHANNEL_ID")
	}

	// ... file upload code ...
}
```

- Loads environment variables.
- Validates required tokens and channel ID.
- Performs the file upload.

### 📝 Usage

- Ensure `SampleFile.md` exists in the project directory.
- Set `SLACK_BOT_TOKEN` and `CHANNEL_ID` in `.env`.
- Run the program to upload the file to Slack.