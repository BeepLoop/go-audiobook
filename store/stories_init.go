package store

import (
	"encoding/json"
	"log"
	"os"

	"github.com/BeepLoop/go-audiobook/types"
)

func InitStoryStore() error {
	_, err := os.Stat("data/stories.json")
	if err == nil {
		log.Println("stories.json already exists")
		return nil
	}
	log.Println("creating stories.json because its missing")

	stories := types.DataStore{
		Stories: []types.Story{},
	}

	data, err := json.Marshal(stories)
	if err != nil {
		return err
	}

	err = os.WriteFile("data/stories.json", data, 0666)
	if err != nil {
		return err
	}

	log.Println("stories.json created successfully")
	return nil
}
