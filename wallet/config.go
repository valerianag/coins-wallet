package wallet

import "github.com/vrischmann/envconfig"

type AppConfig struct {
	PostgresDSN string `envconfig:"POSTGRES_DSN,default=postgresql://postgres:postgres@localhost:5432/wallet_db?sslmode=disable"`
	AppPort     string `envconfig:"APP_PORT,default=:8080"`
}

func NewAppConfig() (AppConfig, error) {
	config := AppConfig{}

	err := envconfig.InitWithOptions(&config, envconfig.Options{AllOptional: true})
	if err != nil {
		return config, err
	}

	return config, nil
}
