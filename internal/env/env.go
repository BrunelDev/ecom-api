package env

import (
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string)string {
	godotenv.Load()
	val := os.
}
