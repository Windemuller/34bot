package main

import (
	"flag"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
var session *discordgo.Session

func init() {
	flag.Parse()
}

func init() {
	var err error
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	session, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func main() {
}
