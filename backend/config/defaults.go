package config

import "github.com/spf13/viper"

// RegisterDefaults for viper
func RegisterDefaults() {
	viper.SetDefault("jwt.token", "")
}
