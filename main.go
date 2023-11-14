package main

import (
	"log"

	"github.com/audiobook/config"
	"github.com/audiobook/server"
	"github.com/audiobook/store"
)

func main() {
	err := config.Init("3000")
	if err != nil {
		log.Fatal(err)
	}

	err = store.InitStorage()
	if err != nil {
		log.Fatal(err)
	}

	server.StartServer(config.Env)
}
