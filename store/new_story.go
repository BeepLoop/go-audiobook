package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/BeepLoop/go-audiobook/types"
)

/*
saves the story to data/stories.json and
automatically reloads the stories after saving
*/
func SaveStoryLocal(story *types.Story) error {
	err := LoadStories()
	if err != nil {
		fmt.Println(err)
		return err
	}

	Stories.Stories = append(Stories.Stories, *story)

	data, err := json.Marshal(Stories)
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.WriteFile("data/stories.json", data, 0666)
	if err != nil {
		log.Println(err)
		return err
	}

	err = LoadStories()
	if err != nil {
		return err
	}

	return nil
}
