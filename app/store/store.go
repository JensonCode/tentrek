package store

import "database/sql"

type Store struct {
	UserStore *UserStore
	AuthStore *AuthStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		UserStore: &UserStore{
			db: db,
		},
		AuthStore: NewAuthStore(),
	}
}
