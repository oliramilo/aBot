package commands

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

/**Bot reply to a message by user**/
func botReply(botid string, session *discordgo.Session, messageSession *discordgo.MessageCreate, availabilities map[string][]int) {
	if botid != messageSession.Author.ID {
		//Reply
		dayOfWeek := int(time.Now().Weekday())

		if dayOfWeek > 5 {
			session.ChannelMessageSend(messageSession.ChannelID, "No one is in the office day")
		} else {
			message := []string{}
			for name, timetable := range availabilities {
				if timetable[dayOfWeek-1] == 1 {
					message = append(message, name+" is in the office today.")
				}
			}
			_, _ = session.ChannelMessageSend(messageSession.ChannelID, strings.Join(message, "\n"))
		}
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
