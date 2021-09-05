package common

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	Prefix = "!"
)

func ProcessMessage(message *discordgo.Message) (bool, string) {
	content := message.Content
	if !strings.HasPrefix(content, Prefix) {
		return false, ""
	}
	if message.Author.Bot {
		return false, ""
	}
	return true, strings.Replace(content, "!", "", 1)
}
