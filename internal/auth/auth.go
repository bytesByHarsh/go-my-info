package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extract API key from the headers
// Example:
// Authorization : ApiKey {insert API Key here}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication Info Found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil

}
