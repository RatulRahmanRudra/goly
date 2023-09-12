package main

import (
	"github.com/RatulRahmanRudra/goly/models"
	"github.com/RatulRahmanRudra/goly/server"
)

func main() {
	models.Setup()
	server.SetupAndListen()
}
