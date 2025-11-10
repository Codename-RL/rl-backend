package config

import (
	"codename-rl/internal/pkg/auth"

	"github.com/spf13/viper"
)

func NewJwt(config *viper.Viper) *auth.JwtService {
	secret := config.GetString("jwt.secret")
	return auth.NewJwtService(secret)
}
