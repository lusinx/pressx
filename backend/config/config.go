package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Init viper
func Init() {
	RegisterDefaults()
	viper.EnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
