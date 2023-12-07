package initializer

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVars() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed loading the env vars ➜ %v", err.Error())
	}
	return nil
}
