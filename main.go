package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Constani/discordbot/commands"
	"github.com/Constani/discordbot/events"
	"github.com/Constani/discordbot/scripts"
	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + scripts.GetBotToken())
	if err != nil {
		fmt.Println("Bot başlatılmadı hata :" + err.Error())
	}
	dg.AddHandler(events.VoiceStateUpdate)
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "" {
		return
	}
	if strings.HasPrefix(m.Content, ".") {
		commands.ExecuteCommand(s, m.Message, ".")
	}
}
