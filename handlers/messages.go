package handlers

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	//если канал == "скриншоты" || "тест-бота"
	if m.ChannelID == "1398324572192899183" || m.ChannelID == "1398301425418899536" {
		s.MessageReactionAdd(m.ChannelID, m.ID, "👍")
		s.MessageReactionAdd(m.ChannelID, m.ID, "👎")

		isBirch, _ := regexp.MatchString(`(?i)бер[её]з(а|ка)`, m.Content)
		isBirchHouse, _ := regexp.MatchString(`(?i)бер[её]зовый дом`, m.Content)

		if isBirch {
			err := s.MessageReactionAdd("1398324572192899183", m.ID, "birch:1421124042232041573")
			if err != nil {
				fmt.Printf("cant add custom reaction to message %s : %s\n", m.ID, err)
				return
			}
		}
		if isBirchHouse {
			err := s.MessageReactionAdd("1398324572192899183", m.ID, "birch_house:1398324572192899183")
			if err != nil {
				fmt.Printf("cant add custom reaction to message %s : %s\n", m.ID, err)
				return
			}
		}
	}
}
