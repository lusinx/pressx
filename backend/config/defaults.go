package config

import "github.com/spf13/viper"

// RegisterDefaults for viper
func RegisterDefaults() {
	viper.SetDefault("jwt.token", "")
	viper.SetDefault("postgres.host", "")
	viper.SetDefault("postgres.port", 0)
	viper.SetDefault("postgres.username", "")
	viper.SetDefault("postgres.password", "")
	viper.SetDefault("postgres.database", "")
}
