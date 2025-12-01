package main

import (
	_ "flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("BOT " + token)
	if err != nil {
		fmt.Println("cant create new session", err)
		return
	}
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildEmojis | discordgo.IntentsMessageContent

	err = dg.Open()
	if err != nil {
		fmt.Println("cant create a connection", err)
		return
	}
	fmt.Println("bot is fully working!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!wl" {
		_, err := s.ChannelMessageSend(m.ChannelID, "test")
		if err != nil {
			fmt.Println("cant send message", err)
			return
		}
	}
}
