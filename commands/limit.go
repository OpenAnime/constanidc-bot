package commands

import (
	"github.com/bwmarrin/discordgo"
)

func HandleLimitCommand(s *discordgo.Session, m *discordgo.Message, game string) {

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Selam",
		Description: game,
	})
}
