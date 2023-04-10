package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost             string `mapstructure:"DB_HOST"`
	DBName             string `mapstructure:"DB_NAME"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBSOURCE           string `mapstructure:"DB_SOURCE"`
	SMTPPORT           string `mapstructure:"SMTP_PORT"`
	SMTPHOST           string `mapstructure:"SMTP_HOST"`
	SMTPPASSWORD       string `mapstructure:"SMTP_PASSWORD"`
	SMTPUSERNAME       string `mapstructure:"SMTP_USERNAME"`
	OauthStateString   string `mapstructure:"OauthStateString"`
	ClientID           string `mapstructure:"ClientID"`
	ClientSecret       string `mapstructure:"ClientSecret"`
	TWAccountSID       string `mapstructure:"ACCOUNT_SID"`
	TWVerifyServiseSID string `mapstructure:"VERIFY_SERVICE_SID"`
	TWAuthTocken       string `mapstructure:"AUTH_TOKEN"`
	TWFromPhone        string `mapstructure:"FROM_PHONE"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "DB_SOURCE", "SMTP_PORT", "SMTP_HOST", "SMTP_PASSWORD", "SMTP_USERNAME", "OauthStateString", "ClientID", "ClientSecret", "ACCOUNT_SID", "VERIFY_SERVICE_SID", "AUTH_TOKEN", "FROM_PHONE",
}

func LoadConfig() (Config, error) {
	var config Config
	fmt.Printf("\n\n\ncfg : %v\n\n\n", config)

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
	fmt.Printf("\n\nconfig : %v\n\n", config)
	return config, nil
}
