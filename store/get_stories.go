package store

import "github.com/audiobook/types"

func GetStories() ([]types.Story, error) {
	var stories []types.Story

	query := "SELECT id, title, author, thumbnail FROM story"
	err := Db_Conn.Select(&stories, query)
	if err != nil {
		return nil, err
	}

	return stories, nil
}
