package bot

import (
	"log"
	"os"

	"github.com/seanstoves/szdiscordbot/commands"

	"github.com/bwmarrin/discordgo"
)

// Start initializes and starts the bot
func Start() error {
	token := os.Getenv("DISCORD_BOT_TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	dg.AddHandler(commands.MessageCreate)

	err = dg.Open()
	if err != nil {
		return err
	}

	// Keep the program running until an interrupt signal is received
	log.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
