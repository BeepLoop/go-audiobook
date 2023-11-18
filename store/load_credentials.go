package store

import (
	"encoding/json"
	"os"

	"github.com/BeepLoop/go-audiobook/types"
)

var Credential types.Admin

func LoadCredential() error {
	data, err := os.ReadFile("data/admin.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Credential)
	if err != nil {
		return err
	}

	return nil
}
