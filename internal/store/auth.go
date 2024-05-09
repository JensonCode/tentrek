package store

import (
	"errors"
	"sync"
	"time"

	"github.com/JensonCode/tentrek/internal/model"
	"github.com/google/uuid"
)

type AuthStore struct {
	OTP  map[string]OTPMapValue
	Lock sync.Mutex
}

type OTPMapValue struct {
	otp string
	model.CreateUserRequest
}

func NewAuthStore() *AuthStore {
	return &AuthStore{
		OTP: make(map[string]OTPMapValue),
	}
}

func (s *AuthStore) StoreOTP(req model.CreateUserRequest, otp string) string {

	uuid := uuid.New().String()
	value := &OTPMapValue{
		otp,
		model.CreateUserRequest{
			Email:    req.Email,
			Password: req.Password,
			Provider: req.Provider,
		},
	}

	s.Lock.Lock()
	s.OTP[uuid] = *value
	s.Lock.Unlock()

	go func() {
		<-time.After(120 * time.Second)
		s.Lock.Lock()
		delete(s.OTP, uuid)
		s.Lock.Unlock()

	}()

	return uuid

}

func (s *AuthStore) DeleteOTP(req *model.EmailVerificationRequest) (*model.CreateUserRequest, error) {

	s.Lock.Lock()
	storedInfo, ok := s.OTP[req.RegisterID]
	if storedInfo.otp == req.OTP {
		delete(s.OTP, req.RegisterID)
	} else {
		ok = false
	}
	s.Lock.Unlock()

	if !ok {
		return nil, errors.New("invaild OTP")
	}

	return &storedInfo.CreateUserRequest, nil
}
