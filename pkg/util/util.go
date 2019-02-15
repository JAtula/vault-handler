package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readToken(token bool) (string, error) {
	if token {
		return "Don't read default token", nil
	}
	data, err := ioutil.ReadFile("./test_token")
	if err != nil {
		return "", err
	}
	err = os.Setenv("VAULT_TOKEN", string(data))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	cb, s := os.LookupEnv("VAULT_TOKEN")
	if s == false {
		return "Couldn't find env.", err
	}
	return cb, nil
}
