package store

import "database/sql"

type Store struct {
	UserStore *UserStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		UserStore: &UserStore{
			db: db,
		},
	}
}
