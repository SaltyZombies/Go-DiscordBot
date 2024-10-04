package main

import (
	"log"

	"github.com/seanstoves/szdiscordbot/bot"
)

func main() {
	// Initialize the bot
	err := bot.Start()
	if err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
}
