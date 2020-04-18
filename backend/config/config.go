package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Init viper
func Init() {
	RegisterDefaults()
	viper.EnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.Set(
		"postgres.config",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			viper.GetString("postgres.host"),
			viper.GetInt("postgres.port"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DATABASE"),
		),
	)
	fmt.Println(viper.GetString("postgres.config"))

}
