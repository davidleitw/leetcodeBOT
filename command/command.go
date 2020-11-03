package command

import "github.com/bwmarrin/discordgo"

var cmd map[string]func(*discordgo.Message)
