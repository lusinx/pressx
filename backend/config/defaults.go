package config

import "github.com/spf13/viper"

// RegisterDefaults for viper
func RegisterDefaults() {
	viper.SetDefault("jwt.token", "")
	viper.SetDefault("postgres.host", "172.16.238.13")
	viper.SetDefault("postgres.port", 5432)
	viper.SetDefault("postgres.username", "")
	viper.SetDefault("postgres.password", "")
	viper.SetDefault("postgres.database", "")
}
