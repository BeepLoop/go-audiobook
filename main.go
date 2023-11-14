package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

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

	OpenBrowser()

	server.InitServer()
	err = server.StartServer(config.Env)
	if err != nil {
		log.Fatal(err)
	}
}

func OpenBrowser() {
	url := fmt.Sprintf("http://%v", config.Env.ListenAddr)

	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}
