package sqllite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
)

func CreateDatabase(sqlFile, dbLoc string) error {
	data, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return err
	}
	database, err := sql.Open("sqlite3", dbLoc)
	if err != nil {
		return err
	}
	statement, err := database.Prepare(string(data))
	statement.Exec()
	return nil
}
