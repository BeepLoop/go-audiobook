package store

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db_Conn *sqlx.DB

func Init() error {
	db, err := sqlx.Open("mysql", "root:Password_1@tcp/audiobook")
	if err != nil {
		return err
	}

	Db_Conn = db

	return nil
}
