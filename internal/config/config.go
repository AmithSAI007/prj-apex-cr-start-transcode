package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv     string `mapstructure:"APP_ENV"`
	HttpPort   string `mapstructure:"HTTP_PORT"`
	ProjectID  string `mapstructure:"PROJECT_ID"`
	Location   string `mapstructure:"LOCATION"`
	TemplateID string `mapstructure:"TEMPLATE_ID"`
	OutputURI  string `mapstructure:"OUTPUT_URI"`
}

func LoadConfig(path string) (*Config, error) {

	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("HTTP_PORT", "8080")
	viper.SetDefault("PROJECT_ID", "your-gcp-project-id")
	viper.SetDefault("LOCATION", "us-central1")
	viper.SetDefault("TEMPLATE_ID", "your-template-id")
	viper.SetDefault("OUTPUT_URI", "info")

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %w", err)
	}

	return &config, nil
}
