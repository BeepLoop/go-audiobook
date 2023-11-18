package store

import (
	"encoding/json"
	"log"
	"os"

	"github.com/BeepLoop/go-audiobook/types"
)

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
