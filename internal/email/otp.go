package email

import (
	"fmt"
	"net/smtp"

	"github.com/JensonCode/tentrek/configs"
	"github.com/JensonCode/tentrek/helpers"
)

func SendOTP(receiver string) (string, error) {
	conf := configs.NewSMTPConf()

	emailAuth := smtp.PlainAuth("", conf.Sender, conf.Password, conf.Host)

	addr := conf.Host + ":" + conf.Port

	to := []string{
		receiver,
	}

	otp := helpers.GenerateOTP()

	msg := fmt.Sprintf("Subject: Tentrek email verification OTP\r\n\r\nyour verification code: %s", otp)

	err := smtp.SendMail(addr, emailAuth, conf.Sender, to, []byte(msg))

	return otp, err
}
