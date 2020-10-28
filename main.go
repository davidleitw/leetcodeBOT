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
	fullID := msg.Author.Username + "#" + msg.Author.Discriminator
	prefix := msg.Content[:1]
	log.Println("prefix = ", prefix)

	switch {
	case msg.Content == "ping":
		_, _ = dis.ChannelMessageSend(msg.ChannelID, "Pong!")
	case msg.Content == "pong":
		_, _ = dis.ChannelMessageSend(msg.ChannelID, "Ping!")
	case msg.Content == "test":
		//_, _ = dis.ChannelMessageSend(msg.ChannelID, fullID+"  ㄐㄐ")
		_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, &discordgo.MessageEmbed{
			Color: 0xFFD700,
			Type:  discordgo.EmbedTypeRich,
			Author: &discordgo.MessageEmbedAuthor{
				URL:     "https://home.gamer.com.tw/homeindex.php?owner=bahamut000",
				Name:    "薯條醬",
				IconURL: "https://cdn.icon-icons.com/icons2/2389/PNG/512/leetcode_logo_icon_145113.png",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   fullID,
					Value:  "test embed message1",
					Inline: false,
				},
				{
					Name:   fullID,
					Value:  "test embed message2",
					Inline: true,
				},
				{
					Name:   "48",
					Value:  "763",
					Inline: true,
				},
				{
					Name:   "4848",
					Value:  "763763",
					Inline: false,
				},
				{
					Name:   "48481",
					Value:  "7631763",
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "一起刷leetcode 一起進步 (･ω´･ )",
				IconURL: "https://imgur.com/DhY5fKW",
			},
		})
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
