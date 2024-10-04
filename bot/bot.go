package bot

import (
	"log"
	"os"

	"github.com/seanstoves/szdiscordbot/commands"

	"github.com/bwmarrin/discordgo"
)

var (
	Token   = os.Getenv("DISCORD_BOT_TOKEN")
	GuildID = os.Getenv("GUILD_ID")
	Command = &discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Say hello!",
	}
)

// Start initializes and starts the bot
func Start() error {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		return err
	}

	dg.AddHandler(commands.MessageCreate)
	dg.AddHandler(commands.ApplicationCmdHello)

	err = dg.Open()
	if err != nil {
		return err
	}
	defer dg.Close()

	_, err = dg.ApplicationCommandCreate(dg.State.User.ID, GuildID, Command)
	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}

	// Keep the program running until an interrupt signal is received
	log.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
