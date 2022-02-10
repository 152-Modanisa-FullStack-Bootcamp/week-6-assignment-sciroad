package config

import (
	"os"
)

const (
	APP_ENV  = "APP_ENV"
	EnvProd  = "production"
	EnvLocal = "local"
)

var env = GetEnv(APP_ENV, EnvLocal)

func GetEnv(key, def string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return env
}

func Env() string {
	return env
}
