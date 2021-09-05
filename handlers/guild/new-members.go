package guild

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/hasanaburayyan/go-karen/common"
)

var (
	logger = log.Default()
)

var (
	testAcceptChannel = "622218194660360202"
	testMemberRole    = "884124601758666765"
	testStudentRole   = "884124660441153626"
)

var (
	FutcAcceptChannel = "570756043831902269"
	FutcMemberRole    = "571354092010602497"
	FutcStudentRole   = "570790341637308447"
)

func ProcessAcceptTermsChannel(session *discordgo.Session, message *discordgo.MessageCreate) {
	process, content := common.ProcessMessage(message.Message)
	if !process {
		return
	}

	if strings.ToLower(content) == "!acceptrules" {
		err := session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, testMemberRole)
		if err != nil {
			logger.Println(err)
		}
		err = session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, testStudentRole)
		if err != nil {
			logger.Println(err)
		}
	}
}
