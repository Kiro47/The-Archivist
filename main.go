package main

import (
	"Archivist/config"
	"Archivist/discord_bot"
	"fmt"
	"os"
)

func main() {
	err := config.LoadConfig("./config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-10)
	}
	discord_bot.BotInit()
	<-make(chan struct{})
}
