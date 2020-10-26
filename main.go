package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	disToken = os.Getenv("DISCORD_TOKEN")
}

var disToken string

func main() {
	if disToken == "" {
		log.Println("No token provided. Please check your bot token.")
		return
	}

	dis, err := discordgo.New("Bot" + disToken)
	if err != nil {
		log.Println("Error creating Discord session: ", err)
		return
	}

	dis.AddHandler(func() {})
	err = dis.Open()
	if err != nil {
		log.Println("discord bot starting error.")
	}

	dis.Close()
}
