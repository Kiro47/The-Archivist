package event_listener

import (
	"Archivist/config"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	config.ZLog.Debug("Discord Message Received:",
		zap.Any("discord_data", msg),
		)
}

func InitMessages(session *discordgo.Session) {
	session.AddHandler(messageCreate)
}