package config

import (
	"codename-rl/internal/pkg/email"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewEmail(v *viper.Viper, log *logrus.Logger) *email.Client {
	fmt.Println("DEBUG SMTP HOST =", v.GetString("smtp.host"))
	fmt.Println("DEBUG SMTP PORT =", v.GetInt("smtp.port"))
	fmt.Println("DEBUG SMTP EMAIL =", v.GetString("smtp.email"))
	fmt.Println("DEBUG SMTP FROM =", v.GetString("smtp.from"))

	cfg := &email.SMTPConfig{
		Host:     v.GetString("smtp.host"),
		Port:     v.GetInt("smtp.port"),
		Email:    v.GetString("smtp.email"),
		Password: v.GetString("smtp.password"),
		From:     v.GetString("smtp.from"),
	}

	client := email.NewClient(cfg)

	log.Info("Email client initialized")

	return client
}
