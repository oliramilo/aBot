/** Discord Bot written in GoLang**/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var tok string
var id string

func getId() {
	args := os.Args[1:]
	if len(args) == 2 {
		tok = args[0]
		id = args[1]
	} else if len(args) == 1 {
		filename := args[0]
		readFile(filename)
	} else {
		fmt.Println("Program requires bot token and botid")
	}
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	tok = scanner.Text()

	scanner.Scan()
	id = scanner.Text()
}

func sendMessage(session *discordgo.Session, message *discordgo.MessageCreate) {

}

func authMessage(message *discordgo.MessageCreate) (authBot bool) {
	authBot = message.Author.ID != id
	return
}

func main() {
	fmt.Println("Program started...")
	fmt.Println("Start time: " + time.Now().String())

}
