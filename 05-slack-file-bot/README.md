# 05 - Slack File Bot

A Slack bot built with Go that uploads a markdown file to a specified Slack channel. This project uses the official Slack Go SDK.

## Features

- Uploads a local markdown file to Slack
- Uses environment variables for configuration
- Simple command-line execution

## Project Structure

- `main.go`: Main application logic for file upload
- `SampleFile.md`: Sample markdown file to upload
- `go.mod`, `go.sum`: Go module files for dependency management

## How to Run

1. Navigate to the project directory:
   ```sh
   cd "05-slack-file-bot"
   ```
2. Create a `.env` file with your Slack tokens:
   ```
   SLACK_BOT_TOKEN=xoxb-your-bot-token
   CHANNEL_ID=C1234567890
   ```
3. Ensure `SampleFile.md` exists and has content
4. Download dependencies:
   ```sh
   go mod tidy
   ```
5. Run the bot:
   ```sh
   go run main.go
   ```

## Requirements

- Slack app with bot token
- Bot added to the target Slack channel
- Channel ID for upload destination

---

This project is part of my Go learning journey!