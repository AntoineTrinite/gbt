package main

import (
	"fmt"
	"github.com/AntoineTrinite/gbt/bot"
)

func main() {
	fmt.Println("HELLO bot")

	bot.BotToken = ""
	bot.Run()
}
