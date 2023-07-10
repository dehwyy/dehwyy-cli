package utils

import (
	e "github.com/dehwyy/dehwyy-cli/error-handler"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	e.WithFatalString(godotenv.Load(".env"), "Cannot load env")
}
