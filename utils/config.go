package utils

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	AppEnv  string `mapstructure:"APP_ENV"`
	AppHost string `mapstructure:"APP_HOST"`

	PgDBHost string `mapstructure:"PG_DB_HOST"`
	PgDBPort string `mapstructure:"PG_DB_PORT"`
	PgDBName string `mapstructure:"PG_DB_NAME"`
	PgDBUser string `mapstructure:"PG_DB_USER"`
	PgDBPass string `mapstructure:"PG_DB_PASS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.SetDefault("APP_ENV", "dev")
	viper.SetDefault("APP_HOST", "127.0.0.1:3000")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
