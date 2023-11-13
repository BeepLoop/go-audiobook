package store

import (
	"github.com/audiobook/types"
)

func GetStoryData(storyId int) (types.Story, []types.Words, error) {
	var story types.Story
	var words []types.Words

	query := "SELECT * FROM story WHERE id=?"
	err := Db_Conn.Get(&story, query, storyId)
	if err != nil {
		return story, words, err
	}

	query = "SELECT id, word, url FROM word WHERE storyId=?"
	err = Db_Conn.Select(&words, query, storyId)
	if err != nil {
		return story, words, err
	}

	return story, words, nil
}
