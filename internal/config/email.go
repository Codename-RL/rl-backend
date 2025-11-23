package config

import (
	"codename-rl/internal/pkg/email"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewEmail(v *viper.Viper, log *logrus.Logger) *email.Client {
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
