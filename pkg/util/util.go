package util

import (
	"io/ioutil"
)

func readToken(token string) (string, error) {
	if token != "" {
		return "Don't read default token", nil
	}
	data, err := ioutil.ReadFile(token)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
