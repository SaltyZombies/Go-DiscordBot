package bot

import (
	"github.com/bwmarrin/discordgo"
)

// Ready is a function that handles the Ready event
func Ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "Ready to serve!")
}

// MessageCreate handles new messages
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "/ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
