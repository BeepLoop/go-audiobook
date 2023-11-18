package store

import (
	"encoding/json"
	"os"

	"github.com/BeepLoop/go-audiobook/types"
)

var Stories types.DataStore

func LoadStories() error {
	data, err := os.ReadFile("data/stories.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Stories)
	if err != nil {
		return err
	}

	return nil
}
