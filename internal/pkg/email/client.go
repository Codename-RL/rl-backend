package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type Client struct {
	cfg SMTPConfig
}

type SMTPConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
	From     string
}

func NewClient(cfg *SMTPConfig) *Client {
	return &Client{cfg: *cfg}
}

func (c *Client) SendOTP(to, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", c.cfg.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")

	body := fmt.Sprintf("Your OTP code is: %s\n\nThis code will expire shortly.", otp)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(c.cfg.Host, c.cfg.Port, c.cfg.Email, c.cfg.Password)

	return d.DialAndSend(m)
}
