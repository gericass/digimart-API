package Infrastructure

import "database/sql"

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
		// TODO
		_, err := tx.Exec("INSERT INTO instruments(test,test2) VALUES (?,?)", 12, 12)
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
