package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key,  fallback string)string {
	godotenv.Load()
	val := os.Getenv(key)
	if val == ""{
	return val
}
return fallback
}
