package message

import (
	"encoding/json"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/hasanaburayyan/go-karen/common"
)

var (
	logger = log.Default()
)

type LogMessage struct {
	Author  string `json:author`
	Bot     bool   `json:bot`
	GuildID string `json:guildId`
	Content string `json:content`
}

func HelpMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	process, content := common.ProcessMessage(message.Message)
	if !process {
		return
	}
	if content != "help" {
		return
	}
	channel := message.ChannelID
	helpMessage := discordgo.MessageEmbed{
		Title:       "Available Commands",
		Description: "Hi There! The following are my available help commands, feel free to try them out!",
		Color:       5763719,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Im Back Baby! This Time With GOPHERS!!!",
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Like this bot? Please leave a ⭐",
			URL:  "https://github.com/hasanaburayyan/go-karen",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "!acceptrules:",
				Value: "if you have not yet accepted the code of conduct rules you will need to use this command in accept-conduct channel before you can access the server",
			},
			{
				Name:  "!recommend:",
				Value: "use this command followed by any recomendation to give us advice and feedback on anything or recomendation on anything like the Server/Karen/Events",
			},
		},
	}
	_, err := session.ChannelMessageSendEmbed(channel, &helpMessage)
	if err != nil {
		logger.Println(err)
	}
}

func LogMessages(session *discordgo.Session, message *discordgo.MessageCreate) {
	var messageLog = LogMessage{
		Author:  message.Author.Username,
		Bot:     message.Author.Bot,
		GuildID: message.GuildID,
		Content: message.Content,
	}

	messageLogOutput, err := json.Marshal(messageLog)
	if err != nil {
		logger.Printf("Could Not Marshal Message %v", message.ID)
	}

	logger.Printf("Message Created! \n%v", string(messageLogOutput))
}

func RepeatMessageBack(session *discordgo.Session, message *discordgo.MessageCreate) {
	process, content := common.ProcessMessage(message.Message)
	if !process {
		return
	}
	channel := message.ChannelID
	session.ChannelMessageSend(channel, content)
}
