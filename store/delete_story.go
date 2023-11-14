package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/audiobook/types"
)

/*
Makes a copy of the stories but without the story with given id,
then reassigns the loaded Stories to the copied stories.

Automatically reloads the stories list after saving the new stories
to stories.json.
*/
func DeleteStoryLocal(id string) error {
	err := LoadStories()
	if err != nil {
		fmt.Println(err)
		return err
	}

	newStories := []types.Story{}

	for _, story := range Stories.Stories {
		if story.Id != id {
			newStories = append(newStories, story)
		}
	}

	Stories.Stories = newStories

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
