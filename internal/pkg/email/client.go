package email

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type Client struct {
	viper  *viper.Viper
	logger *logrus.Logger
}

func (c *Client) NewClient(viper *viper.Viper, logger *logrus.Logger) *Client {
	return &Client{
		viper:  viper,
		logger: logger,
	}
}

func (c *Client) SendOTP(to, otp string) error {
	host := c.viper.GetString("smtp.host")
	port := c.viper.GetInt("smtp.port")
	email := c.viper.GetString("smtp.email")
	password := c.viper.GetString("smtp.password")
	from := c.viper.GetString("smtp.from")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")

	body := fmt.Sprintf("Your OTP code is: %s\n\nThis code will expire in a few minutes.", otp)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(host, port, email, password)

	return d.DialAndSend(m)
}
