# 04 - Slack Age Bot

This Go program implements a Slack bot that calculates a user's age based on their year of birth. It uses the Slacker library for easy Slack bot development.

### 🔧 Imports

```go
import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)
```

- `context`, `fmt`, `log`: Standard libraries for logging and output.
- `os`, `strconv`: For environment variables and string conversion.
- `time`: For getting current year.
- `github.com/joho/godotenv`: For loading environment variables from .env file.
- `github.com/shomali11/slacker`: Library for building Slack bots.

### 🧩 Main Functionality

#### Command Definition

```go
bot.Command("my yob is <year>", &slacker.CommandDefinition{
	Description: "yob calculator",
	Examples:    []string{"my yob is 2025"},
	Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
		year := request.Param("year")

		yob, err := strconv.Atoi(year)
		if err != nil {
			response.Reply("Invalid year. Use a numeric value like 2001.")
			return
		}

		currentYear := time.Now().Year()
		age := currentYear - yob

		if age < 0 {
			response.Reply("Year of birth can not be in the future.")
			return
		}

		response.Reply(fmt.Sprintf("Your age is %d years old", age))
	},
})
```

- Defines a Slack command "my yob is <year>" where <year> is a parameter.
- Parses the year, calculates age, and replies with the age.
- Handles invalid input and future years.

#### Event Printing

```go
func printCommandEvents(events <-chan *slacker.CommandEvent) {
	for event := range events {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
```

- Goroutine to print command events for debugging.

### 🚀 Main Function

```go
func main() {
	godotenv.Load()

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Listen(context.Background())
}
```

- Loads environment variables.
- Creates a new Slacker client with bot and app tokens.
- Starts listening for commands.

### 📝 Usage

- In Slack, type: `my yob is 1990`
- Bot replies: `Your age is 34 years old` (assuming current year is 2024)