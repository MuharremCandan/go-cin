package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Port     string         `mapstructure:"PORT"`
	Database DatabaseConfig `mapstructure:"DATABASE"`
	Cloud    Cloud          `mapstructure:"CLOUD"`
}

type DatabaseConfig struct {
	DbHost string `mapstructure:"DB_HOST"`
	DbPort string `mapstructure:"DB_PORT"`
	DbUser string `mapstructure:"DB_USER"`
	DbPass string `mapstructure:"DB_PASS"`
	DbName string `mapstructure:"DB_NAME"`
}

type Cloud struct {
	CloudinaryURL string `mapstructure:"CLOUDINARY_URL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
