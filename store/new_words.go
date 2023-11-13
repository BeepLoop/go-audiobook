package store

import (
	"log"

	"github.com/audiobook/types"
)

func SaveWords(words []*types.Words, storyId int64) error {
	query := `
            INSERT INTO 
                word (word, url, storyId)
            VALUES (?, ?, ?)
            `

	tx, err := Db_Conn.Beginx()
	if err != nil {
		return err
	}

	for _, word := range words {
		log.Println(word)
		_, err := tx.Exec(query, word.Word, word.Url, storyId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
