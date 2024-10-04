package commands

import "github.com/bwmarrin/discordgo"

// MessageCreate handles new messages
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "/hello" {
		s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Username+"!")
	}
}

func ApplicationCmdHello(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "hello" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hello, " + i.Member.User.Username + "!",
			},
		})
	}
}
