package Infrastructure

import "database/sql"

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
	err = txHandler(tx, f)
	if err != nil {
		return err
	}
	return nil
}
