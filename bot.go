package main

import (
	"encoding/json"
	"fmt"
	"github.com/gempir/go-twitch-irc/v2"
	"os"
)


type Configuration struct {
	Channel   string
	User string
	Oauth  string
}

var configuration = Configuration{}

func main() {


	responseCommands := map[string]func(*twitch.PrivateMessage, *twitch.Client){
		"!tournament": Tournament,
	}

	file, _ := os.Open("conf.json")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("error:", err)
	}

	channel := configuration.Channel
	user := configuration.User
	oauth := configuration.Oauth

	client := twitch.NewClient(user, oauth)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		if message.Message[0] == '!'{
			response := responseCommands[message.Message]
			if response != nil {
				response(&message, client)
			}  else {
				client.Say(channel, "Command not found")
			}
		}


	})

	client.OnConnect(func() {
		client.Say(channel, "bot joined")
	})

	client.Join(channel)

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}

