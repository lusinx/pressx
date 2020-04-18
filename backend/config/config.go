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
			viper.GetString("postgres.username"),
			viper.GetString("postgres.password"),
			viper.GetString("postgres.database"),
		),
	)
	fmt.Println(viper.GetString("postgres.config"), os.Getenv("POSTGRES_HOST"), viper.GetString("postgres.host"))

}
