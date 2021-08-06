package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var BotID string

/**Bot token taken from args**/
func getId() string {
	args := os.Args[1:]
	if len(args) == 1 {
		return args[0]
	} else {
		fmt.Println("Program requires bot token")
	}
	return ""
}

func main() {
	timeStart := time.Now()
	fmt.Println("Program started...")
	fmt.Println("Start time: " + timeStart.String())
	botToken := getId()

	aBot, err := discordgo.New("Bot " + botToken)

	if err != nil {
		fmt.Println(err.Error())
	}

	bot, err := aBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = bot.ID

	err = aBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is Live and Running...")

	<-make(chan struct{})
	return

}

func getDay() int {
	return time.Now().Day()
}
