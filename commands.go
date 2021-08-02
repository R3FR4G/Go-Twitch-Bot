package main

import (
	"github.com/gempir/go-twitch-irc/v2"
)

func Tournament(message *twitch.PrivateMessage, client *twitch.Client) {
	client.Say(configuration.Channel, "The Sonic Speedrunning Community are hosting a Sonic Any% Tournament for SRB2! Check out the details & sign up here: shorturl.at/sGIT8")
}