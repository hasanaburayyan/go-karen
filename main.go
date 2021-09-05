package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/hasanaburayyan/go-karen/handlers/guild"
	"github.com/hasanaburayyan/go-karen/handlers/message"
)

// Bot parameters
var (
	BotToken = flag.String("token", "", "Bot access token")
)

func init() {
	flag.Parse()
}

var session *discordgo.Session

func init() {
	var err error
	session, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatal(err)
	}
	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Defer the closing of session until main function completes
	defer session.Close()

	// Add Handlers To Session
	session.AddHandler(message.LogMessages)
	session.AddHandler(message.HelpMessage)
	session.AddHandler(guild.ProcessAcceptTermsChannel)

	// Keep session open until interruption
	fmt.Println("Bot is now running. Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
