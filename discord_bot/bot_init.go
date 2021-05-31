package discord_bot

import (
	"Archivist/config"
	"Archivist/discord_bot/event_listener"
	"github.com/bwmarrin/discordgo"
	"os"
)

func loadEvents(session *discordgo.Session) {
	event_listener.InitMessages(session)
}
func BotInit() {
	botSession, err := discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		config.ZLog.Panic("Failed to create bot instance.\n" + err.Error())
		os.Exit(-1)
	}
	botUser, err := botSession.User("@me")

	if err != nil {
		config.ZLog.Panic("Failed to obtain Discord BotID.\n" + err.Error())
		os.Exit(-2)
	}

	// Register all Event listeners
	loadEvents(botSession)

	// Open session for bot
	err = botSession.Open()
	if err != nil {
		config.ZLog.Panic("Failed to connect to Discord API.\n" + err.Error())
		os.Exit(-3)
	}
	config.ZLog.Info("Connected to Discord API with BotID " + botUser.ID)

	return
}
