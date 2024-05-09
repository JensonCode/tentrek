package configs

import "os"

type SMTPConf struct {
	Host     string
	Port     string
	Sender   string
	Password string
}

func NewSMTPConf() *SMTPConf {
	return &SMTPConf{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Sender:   os.Getenv("SMTP_SENDER"),
		Password: os.Getenv("SMTP_SENDER_PASSWORD"),
	}
}
