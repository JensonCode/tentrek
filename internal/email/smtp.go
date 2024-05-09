package email

import (
	"fmt"
	"net/smtp"

	"github.com/JensonCode/tentrek/internal/auth"
	"github.com/JensonCode/tentrek/internal/configs"
)

func SendOTP(receiver string) (string, error) {
	conf := configs.NewSMTPConf()

	emailAuth := smtp.PlainAuth("", conf.Sender, conf.Password, conf.Host)

	addr := conf.Host + ":" + conf.Port

	to := []string{
		receiver,
	}

	otp := auth.GenerateOTP()

	msg := fmt.Sprintf("Subject: Tentrek email verification OTP\r\n\r\nyour verification code: %s", otp)

	err := smtp.SendMail(addr, emailAuth, conf.Sender, to, []byte(msg))

	return otp, err
}
