package store

import (
	"log"

	"github.com/audiobook/types"
)

/*
* saves the story to database. returns -1 if err
 */
func SaveStory(story *types.Story) (int64, error) {
	storyQuery := `
        INSERT INTO 
            story (title, author, content, audio, thumbnail) 
        VALUES (?,?,?,?,?)`

	res, err := Db_Conn.Exec(storyQuery, story.Title, story.Author, story.Story, story.AudioUrl, story.ThumbnailUrl)
    if err != nil {
        return -1, err
    }

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, err
	}

	return id, nil
}
