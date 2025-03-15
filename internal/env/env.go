package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valasInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return valasInt
}

func GetBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	value, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}
	return value
}
