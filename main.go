/** Discord Bot written in GoLang**/

package main

import (
	"fmt"
	"time"
	"container/list"
	"github.com/bwmarrin/discordgo"
)

var tok string
var bot_id string
var commands *list

func createSession() {
	timeStart := time.Now()
	fmt.println("Time start: " + timeStart)
	discord, err := discordgo.New("Bot " + "authetication token")
	fmt.print(err)
	fmt.print(discord)
}

func processCommand(session *discordgo.Session, message *discordgo.Mes)

func main() {
	createSession()
	
}
