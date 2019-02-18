package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

//ReadTokenFromFile and inject into env variable
func ReadTokenFromFile(filename string) (string, error) {
	cb, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(cb), nil
}

//SetVaultToken to environment variable VAULT_TOKEN
func SetVaultToken(token string) error {
	err := os.Setenv("VAULT_TOKEN", token)
	if err != nil {
		return err
	}
	return nil
}

// WriteSecretPayload to a .json dump
func WriteSecretPayload(path string, data string) (string, error) {
	payload := []byte(data)
	err := ioutil.WriteFile(path, payload, 0644)
	if err != nil {
		return "", err
	}
	response := fmt.Sprintf("Secret written to %s", path)
	return response, nil
}
