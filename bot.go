package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
)


type Configuration struct {
	Channel   string
	Oauth  string
}

func main() {


	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	channel := configuration.Channel
	oauth := configuration.Oauth

	client := twitch.NewClient(channel, oauth)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
		if message.Message == "!tournament" {
			client.Say(channel, "The Sonic Speedrunning Community are hosting a Sonic Any% Tournament for SRB2! Check out the details & sign up here: shorturl.at/sGIT8")
		}
	})

	client.OnConnect(func() {
		client.Say("leminn", "pepegabot initalized")
	})

	client.Join(	"leminn")

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
