package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// NewViper is a function to load config from config.json
// You can change the implementation, for example load from env file, consul, etcd, etc
func NewViper() *viper.Viper {
	_ = godotenv.Load()
	// ---------------------------
	// Load config.json
	// ---------------------------
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error loading config.json: %w", err))
	}

	// ---------------------------
	// Load .env
	// ---------------------------
	//env := viper.New()
	//env.SetConfigFile(".env")
	//env.SetConfigType("env")
	//env.AddConfigPath(".")
	//
	//if err := env.ReadInConfig(); err == nil {
	//	// merge .env INTO config.json (override json values)
	//	if err := config.MergeConfigMap(env.AllSettings()); err != nil {
	//		panic(fmt.Errorf("fatal merging config file: %w", err))
	//	}
	//}

	// also read OS env vars
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	return config
}
