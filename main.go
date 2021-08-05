/** Discord Bot written in GoLang**/

package main


import (
	"fmt"
	"os"
	"time"
	"container/list"
	"github.com/bwmarrin/discordgo"
)



var {
	Token = ""
	BotID string
}


func processCommand(session *discordgo.Session,message *discordgo.Mes)

func main() {
	timeStart = time.Now()

	commandList := list.New()

	discord , err := discordgo.New("Bot " + "authentication token")

}