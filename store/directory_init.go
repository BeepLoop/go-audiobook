package store

import (
	"log"
	"os"
)

func InitDirectories() error {
	dataDirectories := []string{"data/", "assets/audios/", "assets/thumbnails/", "assets/downloads/"}

	for _, dir := range dataDirectories {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Printf("Successfully created path: %s", dir)
	}

	return nil
}
