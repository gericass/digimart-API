package Infrastructure

import "database/sql"

func TXHandler(db sql.DB, f func(tx sql.Tx)) error {
	return nil
}

func (inst *Instrument) Insert(db sql.DB) error {
	var err error
	f := func(tx sql.Tx) {
		tx.Exec("INSERT INTO instruments(test,test2) VALUES (?,?)", 12, 12)
	}
	err = TXHandler(db,f)
	if err != nil {
		return err
	}
	return nil
}
