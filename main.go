package main

import (
	_ "flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/snrKostik/birch-bot/handlers"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func goEnvVar(key string) string {
	err := godotenv.Load(key)
	if err != nil {
		fmt.Printf("cant load key %s : %s", key, err)
	}

	return key
}
func initBot(token string) *discordgo.Session {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Printf("cant create session: %s", err)
	}
	return dg
}

func main() {
	token := goEnvVar("DISCORD_BOT_TOKEN")
	appID := goEnvVar("DISCORD_APP_ID")

	dg := initBot(token)

	dg.AddHandler(handlers.MessageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildEmojis | discordgo.IntentsMessageContent
	handlers.ChangeStatus(dg, appID)

	open(dg)

	fmt.Println("bot is fully working!")

	close(dg)
}

func open(dg *discordgo.Session) {
	err := dg.Open()
	if err != nil {
		fmt.Println("cant create connection: ", err)
		return
	}
}

func close(dg *discordgo.Session) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
