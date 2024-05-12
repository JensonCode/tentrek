package model

import (
	"time"

	"github.com/JensonCode/tentrek/helpers"
	"github.com/google/uuid"
)

type User struct {
	UID       uuid.UUID `json:"uid"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Provider string `json:"provider"`
	Avatar   string `json:"avatar,omitempty"`
}

func (u *User) Init(req *CreateUserRequest) error {

	uuid := uuid.New()

	if req.Provider != "app" {
		req.Password = helpers.GenerateRandomPassword()
	}

	hashed, err := helpers.HashPassword(req.Password)
	if err != nil {
		return err
	}

	u.UID = uuid
	u.Email = req.Email
	u.Password = hashed
	u.Avatar = req.Avatar
	u.Provider = req.Provider
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return nil
}
