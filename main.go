package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/leetcode"
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
	log.Printf("%s: %s\n", msg.Author, msg.Content)

	if (msg.Author.ID == dis.State.User.ID) || msg.GuildID == "" {
		return
	}

	// 用戶暱稱
	// fullID := msg.Author.Username + "#" + msg.Author.Discriminator
	prefix := msg.Content[:1]
	log.Println("prefix = ", prefix)

	switch {
	case msg.Content == "ping":
		_, _ = dis.ChannelMessageSend(msg.ChannelID, "Pong!")
	case msg.Content == "pong":
		_, _ = dis.ChannelMessageSend(msg.ChannelID, "Ping!")
	case msg.Content == "test":
		var problems []*model.Problem
		_48, _ := model.SearchWithIDTest(48)
		_763, _ := model.SearchWithIDTest(763)
		problems = append(problems, _48)
		problems = append(problems, _763)

		ps := leetcode.ProblemEmbedMessage(problems)
		_, err := dis.ChannelMessageSendEmbed(msg.ChannelID, ps)
		if err != nil {
			log.Println("err = ", err)
		}

	case prefix == "!":
		command := msg.Content[1:]
		log.Println("In command area, command = ", command)
		problem := regexp.MustCompile("^problem")
		switch {
		case problem.MatchString(command):
			field := strings.Fields(command)

			num, err := strconv.Atoi(field[1])
			if err != nil {
				_, _ = dis.ChannelMessageSend(msg.ChannelID, "Please input a problem number.")
				return
			}

			p, err := model.SearchWithIDTest(num)
			if err != nil {
				_, _ = dis.ChannelMessageSend(msg.ChannelID, err.Error())
				return
			}
			_, _ = dis.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%d: %s => %s\n", p.ProblemID, p.ProblemTitle, p.ProblemURL))
		}
	}

}
