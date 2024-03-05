package auth

import (
	"errors"
	"net/http"
)

func GetApiKey(headers http.Header) (string, error) {

	value := headers.Get("Authorization")
	if value == "" {
		return "", errors.New("invalid auth key provided")
	}

	return value, nil
}
