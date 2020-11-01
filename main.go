package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/bot"
	"github.com/davidleitw/leetcodeBOT/model"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	disToken = os.Getenv("DISCORD_TOKEN")
	log.Println(disToken)
}

var disToken string

func main() {
	// model.CreateLeetCodeProblemsTable()
	if disToken == "" {
		log.Println("No token provided. Please check your discord bot token.")
		return
	}

	discord, err := discordgo.New("Bot " + disToken)
	if err != nil {
		log.Println("Error creating Discord session: ", err)
		return
	}
	defer discord.Close()

	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)
	// discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	discord.AddHandler(ready)
	discord.AddHandler(messageHandler)
	// discordgo.IntentsAll

	err = discord.Open()
	if err != nil {
		log.Println("discord bot starting error = ", err)
		return
	}

	fmt.Println("bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

// 啟動discord bot
func ready(s *discordgo.Session, event *discordgo.Ready) {
	_ = s.UpdateStatus(0, "leetcode helping bot")
}

func messageHandler(dis *discordgo.Session, msg *discordgo.MessageCreate) {
	log.Println(msg.Author.ID, ": ", msg.Content)
	if (msg.Author.ID == dis.State.User.ID) || msg.GuildID == "" || msg.Author.ID == "235088799074484224" {
		return
	}

	// 用戶暱稱
	// fullID := msg.Author.Username + "#" + msg.Author.Discriminator

	switch {
	case strings.Contains(msg.Content, "ㄐㄐ"):
		_, _ = dis.ChannelMessageSend(msg.ChannelID, "ㄐㄐ")
	case msg.Content == "ping":
		_, _ = dis.ChannelMessageSend(msg.ChannelID, "Pong!")
	case msg.Content == "pong":
		_, _ = dis.ChannelMessageSend(msg.ChannelID, "Ping!")
	case msg.Content == "test":
		var problems []*model.Problem
		_48, _ := model.SearchWithProblemIDTest(48)
		_763, _ := model.SearchWithProblemIDTest(763)
		problems = append(problems, _48)
		problems = append(problems, _763)

		ps := bot.ProblemsEmbedMessage(problems)
		_, err := dis.ChannelMessageSendEmbed(msg.ChannelID, ps)
		if err != nil {
			log.Println("err = ", err)
		}
	}

}
