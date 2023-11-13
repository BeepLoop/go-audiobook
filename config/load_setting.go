package config

import (
	"encoding/json"
	"log"
	"os"
)

func LoadSettings() Setting {
	file, err := os.ReadFile("config/voice_config.json")
	if err != nil {
		log.Fatal("Error loading voice configuration file")
	}

	var setting Setting
	err = json.Unmarshal(file, &setting)
	if err != nil {
		log.Fatal("Please check the voice configuration file")
	}

	return setting
}
