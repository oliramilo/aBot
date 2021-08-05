package commands

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

/**Bot reply to a message by user**/
func botReply(botid string, session *discordgo.Session, message *discordgo.MessageCreate) {
	if botid != message.Author.ID {
		//Reply
	}

}

func getCommands() (commands []string) {
	file, err := os.Open("command_list")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, strings.TrimSpace(scanner.Text()))
	}
	return
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
