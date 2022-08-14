package postgres

import "database/sql"

type Database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

func (d *Database) BeginTransaction() (*Transaction, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}

	return &Transaction{tx: tx}, nil
}

type Transaction struct {
	tx *sql.Tx
}

func (t *Transaction) Commit() error {
	err := t.tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) RollbackFunc(err error) func() error {
	return func() error {
		if err != nil {
			err := t.tx.Rollback()
			if err != nil {
				return err
			}

			return nil
		}
		return nil
	}
}
