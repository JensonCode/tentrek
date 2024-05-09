package store

import (
	"database/sql"
	"errors"

	err "github.com/JensonCode/tentrek/internal/error"
	"github.com/JensonCode/tentrek/internal/model"
)

const (
	ERR_EMAIL_USED  = err.NewError("store user: " + "email has been used")
	ERR_INIT_USER   = err.NewError("store user: " + "initiate new user failed")
	ERR_INSERT_USER = err.NewError("store user: " + "insert user to DB failed")
	ERR_SCAN_USER   = err.NewError("store user: " + "scan user failed")
	ERR_FIND_USER   = err.NewError("store user: " + "user not found")
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
		return nil, errors.New(ERR_EMAIL_USED.Error())
	}

	var newUser = new(model.User)
	err = newUser.Init(req)
	if err != nil {
		return nil, errors.New(ERR_INIT_USER.Error())
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
		return nil, errors.New(ERR_INSERT_USER.Error())
	}

	return newUser, nil
}

func (s *UserStore) IsExist(email string) (bool, error) {
	var count int

	query := `SELECT COUNT(*) FROM users WHERE email = $1`

	err := s.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, errors.New(ERR_SCAN_USER.Error())
	}

	return count > 0, nil
}
