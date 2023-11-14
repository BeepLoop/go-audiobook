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

	admin := types.Admin{
		Username: "admin",
		Password: "admin",
	}

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

func InitDownloadStore() error {
	_, err := os.Stat("assets/downloads/")
	if err != nil {
		log.Println("creating downloads folder because its missing")
		err = os.Mkdir("assets/downloads/", 0777)
		if err != nil {
			log.Println("Could not create folder for downloads")
		}
		log.Println("Dowloads folder created successfully")

		return err
	}

	return nil
}
