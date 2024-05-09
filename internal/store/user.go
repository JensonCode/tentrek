package store

import (
	"database/sql"

	"github.com/JensonCode/tentrek/internal/model"
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) InsertUser(req *model.CreateUserRequest) (*model.User, error) {

	exist, err := s.IsExist(req.Email)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, err
	}

	var newUser = new(model.User)
	err = newUser.Init(req)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO users 
	(uid, email, password, username, avatar, provider, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err = s.db.Exec(query,
		newUser.UID,
		newUser.Email,
		newUser.Password,
		newUser.Username,
		newUser.Avatar,
		newUser.Provider,
		newUser.CreatedAt,
		newUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *UserStore) IsExist(email string) (bool, error) {
	var count int

	query := `SELECT COUNT(*) FROM users WHERE email = $1`

	err := s.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *UserStore) FindByField(field string, value any) (*model.User, error) {

	query := "SELECT * FROM users WHERE " + field + " = $1"

	row := s.db.QueryRow(query, value)

	user, err := scanRow(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func scanRow(row *sql.Row) (*model.User, error) {
	user := new(model.User)

	err := row.Scan(
		&user.UID,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.Avatar,
		&user.Provider,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
