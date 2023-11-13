package server

import (
	"log"

	"github.com/audiobook/config"
)

func StartServer(env *config.Config) {
	InitServer()

	if err := Router.Run(env.ListenAddr); err != nil {
		log.Fatal(err)
	}
}
