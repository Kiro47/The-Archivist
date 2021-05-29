package main

import (
	"Archivist/config"
	"Archivist/discord_bot"
	"log"
)

func main() {
	err := config.LoadConfig("./config.yaml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	config.InitLogger()

	discord_bot.BotInit()
	<-make(chan struct{})
}
