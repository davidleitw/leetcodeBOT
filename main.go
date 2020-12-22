package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/davidleitw/leetcodeBOT/model"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/bot"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	disToken = os.Getenv("DISCORD_TOKEN")
}

var disToken string

func main() {
	model.ConnectionDatabase()
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

	discord.AddHandler(ready)
	discord.AddHandler(messageHandler)

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
	if (msg.Author.ID == dis.State.User.ID) || msg.GuildID == "" || msg.Author.Bot || len(msg.Content) == 0 {
		return
	}

	// log.Println(msg.Member.Nick) 暱稱
	if strings.Contains(msg.Content, "ㄐㄐ") {
		_, _ = dis.ChannelMessageSend(msg.ChannelID, "ㄐㄐ")
		return
	}

	// fullID := msg.Author.Username + "#" + msg.Author.Discriminator
	if strings.HasPrefix(msg.Content, "$") {
		// cmd := strings.Split(msg.Content, " ")[1:]
		command := strings.Split(msg.Content, " ")
		log.Println("command = ", command)
		if len(command) > 1 {
			cmd := command[1:]
			switch cmd[0] {
			// case "search":
			// 	message, err := bot.Search(cmd)
			// 	if err != nil {
			// 		_, _ = dis.ChannelMessageSend(msg.ChannelID, err.Error())
			// 		return
			// 	}
			// 	_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, message)
			// 	return
			case "--help", "-h", "help":
				help := bot.HelpMessage()
				_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, help)
				return
			case "ls", "list", "List":
				message, err := bot.List(msg, cmd)
				if err != nil {
					_, _ = dis.ChannelMessageSend(msg.ChannelID, err.Error())
					return
				}
				_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, message)
				return
			case "add":
				message, err := bot.Add(msg, cmd)
				if err != nil {
					_, _ = dis.ChannelMessageSend(msg.ChannelID, err.Error())
					return
				}
				_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, message)
				return
			case "set", "Set", "SET":
				message, err := bot.Set(msg.GuildID, cmd)
				if err != nil {
					_, _ = dis.ChannelMessageSend(msg.ChannelID, err.Error())
					return
				}
				_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, message)
			case "rm", "del":
				message, err := bot.Remove(msg, cmd)
				if err != nil {
					_, _ = dis.ChannelMessageSend(msg.ChannelID, err.Error())
					return
				}
				_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, message)
				return
			case "draw", "Draw", "抽":
			case "next", "Next", "n":
				message, err := bot.GetNextStudyGroupInfo(msg.GuildID, cmd)
				if err != nil {
					_, _ = dis.ChannelMessageSend(msg.ChannelID, err.Error())
					return
				}
				_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, message)
				return
			case "clear", "CLEAR", "Clear":
				message, err := bot.Clear(msg, cmd)
				if err != nil {
					_, _ = dis.ChannelMessageSend(msg.ChannelID, err.Error())
					return
				}
				_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, message)
				return
			case "上車", "開車":
				_, _ = dis.ChannelMessageSend(msg.ChannelID, "https://media1.tenor.com/images/1e00408e429e6e101b5193c74f136475/tenor.gif")
			default:
				help := bot.HelpMessage()
				_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, help)
				return
			}
		} else {
			help := bot.HelpMessage()
			_, _ = dis.ChannelMessageSendEmbed(msg.ChannelID, help)
			return
		}
	}
}
