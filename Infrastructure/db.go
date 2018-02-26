package Infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func txHandler(tx *sql.Tx, f func(tx *sql.Tx) error) error {
	var err error
	err = f(tx)
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	return err
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:mysql@tcp(db:3306)/digimart")
	if err != nil {
		return nil, err
	}
	return db, nil
}
