package util

import (
	"os"
	"strconv"
)

func GetEnvOrDefault(key string, or string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return or
}

func GetEnvOrDefaultInt(key string, or int) int {
	if v, e := strconv.Atoi(os.Getenv(key)) ;  e == nil {
		return v
	}
	return or
}