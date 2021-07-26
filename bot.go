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

func main() {


	responseCommands := map[string]string{
		"!tournament": "The Sonic Speedrunning Community are hosting a Sonic Any% Tournament for SRB2! Check out the details & sign up here: shorturl.at/sGIT8",
	}

	file, _ := os.Open("conf.json")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
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
			if response != "" {
				client.Say(channel,response)
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

