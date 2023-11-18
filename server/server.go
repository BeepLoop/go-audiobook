package server

import (
	"log"

	"github.com/BeepLoop/go-audiobook/config"
)

func StartServer(env *config.Config) error {
	err := Router.Run(env.ListenAddr)
	if err != nil {
		log.Println(err)
		return err
	}

    return nil
}
