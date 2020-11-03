package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/davidleitw/leetcodeBOT/bot"
	"github.com/davidleitw/leetcodeBOT/leetcode"
	"github.com/davidleitw/leetcodeBOT/model"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	disToken = os.Getenv("DISCORD_TOKEN")
}

var disToken string

func main() {
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
	log.Println(msg.Author.ID, ": ", msg.Content)
	log.Println("msg = ", msg.Content, ", len of msg is ", len(msg.Content))
	if (msg.Author.ID == dis.State.User.ID) || msg.GuildID == "" || msg.Author.ID == "235088799074484224" || len(msg.Content) == 0 {
		return
	}

	if strings.Contains(msg.Content, "ㄐㄐ") {
		_, _ = dis.ChannelMessageSend(msg.ChannelID, "ㄐㄐ")
		return
	}
	// 用戶暱稱
	// fullID := msg.Author.Username + "#" + msg.Author.Discriminator
	if strings.HasPrefix(msg.Content, "$") {
		cmd := strings.Split(msg.Content, " ")[1:]
		switch cmd[0] {
		case "search":
			if cmd[1] != "-d" && len(cmd) == 2 {
				problemID, err := strconv.Atoi(cmd[1])
				if err != nil {
					_, _ = dis.ChannelMessageSend(msg.ChannelID, "請輸入數字，以便於查詢題目。")
					return
				}

				problem, err := leetcode.SearchProblemWithID(problemID)
				if err != nil {
					_, _ = dis.ChannelMessageSend(msg.ChannelID, "資料庫內查無資料，請確認problem ID是否正確。")
					return
				}

				msgs := bot.ProblemsMsg([]*model.Problem{problem})
				_, err = dis.ChannelMessageSendEmbed(msg.ChannelID, msgs)
				if err != nil {
					log.Println("err = ", err)
					return
				}

			} else if (cmd[1] == "-d" || cmd[1] == "--detail") && len(cmd) == 3 {
				problemID, err := strconv.Atoi(cmd[2])
				if err != nil {
					_, _ = dis.ChannelMessageSend(msg.ChannelID, "請輸入數字，以便於查詢題目。")
					return
				}

				problem, err := leetcode.SearchProblemWithID(problemID)
				if err != nil {
					_, _ = dis.ChannelMessageSend(msg.ChannelID, "資料庫內查無資料，請確認problem ID是否正確。")
					return
				}
				msgs := bot.ProblemsDetailMsg([]*model.Problem{problem})
				_, err = dis.ChannelMessageSendEmbed(msg.ChannelID, msgs)
				if err != nil {
					log.Println("err = ", err)
					return
				}
			} else {
				_, _ = dis.ChannelMessageSend(msg.ChannelID, "message error, please check help message(--help, -h)")
			}
		case "--help", "-h", "help":
			_, _ = dis.ChannelMessageSend(msg.ChannelID, "help message area.")

		case "ls":
		case "add":
		case "set":
		case "rm", "del":
		case "draw", "Draw", "抽":

		default:
			_, _ = dis.ChannelMessageSend(msg.ChannelID, "message error, please check help message(--help, -h)")
			return
		}
	}
}
