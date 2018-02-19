package Infrastructure

import "database/sql"

func TXHandler(db sql.DB, f func(tx sql.Tx)) error {

}

func (inst *Instrument) Insert(db sql.DB) error {
	var err error
	f := func(tx sql.Tx) {
		tx.Exec("INSERT INTO instruments()")
	}


	if err != nil {
		return err
	}
	return nil
}
