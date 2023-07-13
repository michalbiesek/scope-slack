package main

import (
	"os"
)

func main() {

	token := os.Getenv("SLACK_BOT_TOKEN")
	// add error handling
	sender := NewSlackSender(token, "test")
	sender.SendNotification()
}
