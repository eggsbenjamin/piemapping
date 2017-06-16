package repository

import "database/sql"

//	emulates api of native 'database/sql' 'Rows' struct to enable easier mocking
type RowIterator interface {
	Next() bool
	Scan(...interface{}) error
	Close() error
	Err() error
}

//	emulates subset of api exposed by native 'database/sql' 'DB' struct to enable easier mocking
type Queryer interface {
	Query(string, ...interface{}) (RowIterator, error)
}

//	wrapper for native 'database/sql' DB struct to enable easier mocking
type DBWrapper struct {
	db *sql.DB
	Queryer
}

//	constructor
func NewDBWrapper(db *sql.DB) *DBWrapper {
	return &DBWrapper{
		db: db,
	}
}

func (w *DBWrapper) Query(query string, args ...interface{}) (RowIterator, error) {
	return w.db.Query(query, args...)
}
