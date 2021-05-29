package discord_bot

import (
	"Archivist/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

var BotID string
var BotSession *discordgo.Session

func BotInit() {
	BotSession, err := discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		fmt.Println("Failed to connect Archivist to Discord.\n" + err.Error())
		os.Exit(-1)
	}
	botUser, err := BotSession.User("@me")

	if err != nil {
		fmt.Println("Failed to obtain Discord BotID.\n" + err.Error())
		os.Exit(-2)
	}
	BotID = botUser.ID

	// Open session for bot
	err = BotSession.Open()
	if err != nil {
		fmt.Println("Failed to connect to Discord API.\n" + err.Error())
		os.Exit(-3)
	}
	fmt.Println("Connected to Discord API with BotID " + BotID)

	return
}
