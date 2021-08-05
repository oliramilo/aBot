package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

/**Bot reply to a message by user**/
func botReply(botid string, session *discordgo.Sessions, message *discordgo.MessageCreate) {
	if botid != message.Author.ID {
		//Reply
	}

}

func processCommand(message string, commands []string) {
	toks := strings.Split(message, " ")
	command := strings.ToUpper(toks[0])
	for _, sCommand := range commands {
		if command == sCommand {
			break
		}
	}
}
