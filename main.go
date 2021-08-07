package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var availabilities map[string][]int = getAvailabilities("availability.csv")

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

	if botToken == "" {
		return
	}

	aBot, err := discordgo.New("Bot " + botToken)

	if err != nil {
		fmt.Println(err.Error())
	}

	bot, err := aBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = bot.ID

	aBot.AddHandler(botReply)
	err = aBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is Live and Running...")

	<-make(chan struct{})
	return

}

/**Bot reply to a message by user**/
func botReply(botid string, session *discordgo.Session, messageSession *discordgo.MessageCreate) {

	if botid != messageSession.Author.ID {
		//Reply

		if messageSession.Content == "/today" {
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
		} else {
			fmt.Println("No commands detected")
		}
	}

}

/**Reads file containing the weekly availabilities of members and returns
a map by member name containing an array of integers, 1=available, 0=not available**/
func getAvailabilities(filename string) map[string][]int {
	availabilities := map[string][]int{}
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		toks := strings.Split(line, ",")
		availabilities[toks[0]] = proc(toks, 1, len(toks))
	}
	return availabilities
}

/**Return integer array of daily availabilities **/
func proc(arr []string, lowb int, upb int) []int {
	nums := []int{}
	for i := lowb; i < upb; i++ {
		numConv, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		nums = append(nums, numConv)
	}
	return nums
}
