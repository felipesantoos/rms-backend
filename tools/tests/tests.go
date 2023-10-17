package tests

import (
	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	godotenv.Load("../../../ui/api/app/.env-for-local")
}
