package initializers

import (
	"github.com/joho/godotenv"
	"github.com/vineet12344/jwtgolang/errorhandlers"
)

func LoadEnv() {
	err := godotenv.Load()
	errorhandlers.CheckError(err)
}
