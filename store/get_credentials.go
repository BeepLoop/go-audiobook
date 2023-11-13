package store

import "github.com/audiobook/types"

func GetCredentials(username string) (*types.User, error) {
	var user types.User

	query := "SELECT username, password FROM user WHERE username=?"
	err := Db_Conn.Get(&user, query, username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
