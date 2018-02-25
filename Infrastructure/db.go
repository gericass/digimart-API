package Infrastructure

import (
	"database/sql"
)

func TXHandler(tx *sql.Tx, f func(tx *sql.Tx) error) error {
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

func (inst *Instrument) Insert(db sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f := func(tx *sql.Tx) error {
		_, err := tx.Exec("INSERT INTO instruments(`name`,`category`,`price`,`condition`,`status`,`url`,`register_date`) VALUES (?,?,?,?,?,?,?)", inst.Name, inst.Category, inst.Price, inst.Condition, inst.Status, inst.URL, inst.RegisterDate)
		if err != nil {
			return err
		}
		return nil
	}
	err = TXHandler(tx, f)
	if err != nil {
		return err
	}
	return nil
}
