package store

func DeleteStory(storyId int) error {
	query := "DELETE FROM story WHERE id=?"

	_, err := Db_Conn.Exec(query, storyId)
	if err != nil {
		return err
	}

	return nil
}
