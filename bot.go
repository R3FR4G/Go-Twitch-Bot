package main

import (
	"encoding/json"
	"fmt"
	"github.com/gempir/go-twitch-irc/v2"
	"math/rand"
	"os"
	"strconv"
)


type Configuration struct {
	Channel   string
	User string
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
	user := configuration.User
	oauth := configuration.Oauth

	client := twitch.NewClient(user, oauth)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
		if message.Message == "!tournament" {
			client.Say(channel, "The Sonic Speedrunning Community are hosting a Sonic Any% Tournament for SRB2! Check out the details & sign up here: shorturl.at/sGIT8")
		} else if message.Message == "!dice" {
			client.Say(channel, "You rolled: " + strconv.Itoa(rand.Intn(6)))
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
