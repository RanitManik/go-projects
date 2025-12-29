# 04 - Slack Age Bot

A Slack bot built with Go that calculates a user's age based on their year of birth. This project uses the Slacker library for easy Slack integration.

## Features

- Responds to Slack commands to calculate age
- Validates input (numeric year, not in future)
- Uses environment variables for tokens

## Project Structure

- `main.go`: Main application logic for the Slack bot
- `go.mod`, `go.sum`: Go module files for dependency management

## How to Run

1. Navigate to the project directory:
   ```sh
   cd "04-slack-age-bot"
   ```
2. Create a `.env` file with your Slack tokens:
   ```
   SLACK_BOT_TOKEN=xoxb-your-bot-token
   SLACK_APP_TOKEN=xapp-your-app-token
   ```
3. Download dependencies:
   ```sh
   go mod tidy
   ```
4. Run the bot:
   ```sh
   go run main.go
   ```

## Usage

In Slack, send a message to the bot:
- `my yob is 1990` - Bot will reply with your age

## Requirements

- Slack app with bot token and app token
- Bot added to a Slack workspace

---

This project is part of my Go learning journey!