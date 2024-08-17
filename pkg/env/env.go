package env

import (
	"os"
)

func GetEnv() string {
	env := os.Getenv("LOANSRV_ENV")
	if env == "" {
		return EnvDevelopment
	}

	return env
}
