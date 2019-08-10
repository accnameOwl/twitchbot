package main

import (
	"github.com/accnameowl/twitchbot/bot"
	"github.com/joho/godotenv"
)

func main() {
	//load environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	bot := bot.New()
	bot.Connect()
	bot.RuntimeQuotes()
}
