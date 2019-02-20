package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func _checkError(e error) {
	if e != nil {
		panic(e)
	}
}

//ReadTokenFromFile and inject into env variable
func ReadTokenFromFile(filename string) (string, error) {
	cb, err := ioutil.ReadFile(filename)
	_checkError(err)
	return string(cb), nil
}

//SetVaultToken to environment variable VAULT_TOKEN
func SetVaultToken(token string) error {
	err := os.Setenv("VAULT_TOKEN", token)
	_checkError(err)
	return nil
}

type secretPayload struct {
	Data map[string]interface{} `json:"data"`
}

// WriteSecretPayload to a .env dump
func WriteSecretPayload(path string, data string) error {
	res := secretPayload{}
	err := json.Unmarshal([]byte(data), &res)
	_checkError(err)
	var str strings.Builder
	for key, value := range res.Data {
		str.WriteString("export ")
		str.WriteString(strings.ToUpper(key))
		str.WriteString("=")
		str.WriteString(value.(string))
		str.WriteString("\n")
	}
	err = ioutil.WriteFile(path, []byte(str.String()), 0644)
	_checkError(err)
	return nil

}
