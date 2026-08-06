package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extracts the API key from the request headers of an http request. It looks for the "Authorization" header and returns the API key if found. If the header is missing or empty, it returns an error.

// example:
// Authorization; ApiKey {insert your api key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing Authorization header")
	}
	vals := strings.SplitN(val, " ", 2)
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("invalid Authorization header format")
	}
	return vals[1], nil
}
