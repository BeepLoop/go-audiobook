package store

import (
	"github.com/audiobook/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db_Conn *sqlx.DB

func Init() error {
	db, err := sqlx.Connect("mysql", config.Env.Dsn)
	if err != nil {
		return err
	}

	Db_Conn = db

	return nil
}
