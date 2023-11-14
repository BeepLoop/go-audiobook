package store

import (
	"encoding/json"
	"log"
	"os"

	"github.com/audiobook/types"
)

func InitLocalStorage() error {
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

func InitAdminStore() error {
	_, err := os.Stat("data/admin.json")
	if err == nil {
		log.Println("admin.json already exists")
		return nil
	}
	log.Println("creating admin.json because its missing")

	admin := types.Admin{}

	data, err := json.Marshal(admin)
	if err != nil {
		return err
	}

	err = os.WriteFile("data/admin.json", data, 0666)
	if err != nil {
		return err
	}

	log.Println("admin.json created successfully")
	return nil
}
