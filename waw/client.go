package waw

import (
	"errors"
	"os"
)

func GetAPIKey(key ...string) (string, error) {
	if len(key) == 0 {
		if envKey := os.Getenv("GOWAW_KEY"); envKey != "" {
			return envKey, nil
		} else {
			return "", errors.New("API Key could not be found in GOWAW_KEY env var or in the parameters")
		}
	} else {
		return key[0], nil
	}
}
