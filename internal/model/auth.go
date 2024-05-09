package model

type LoginReqeust struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailVerificationRequest struct {
	RegisterID string `json:"register_id"`
	OTP        string `json:"otp"`
}
