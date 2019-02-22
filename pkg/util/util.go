package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func _checkError(e error) {
	if e != nil {
		panic(e)
	}
}

//SetVaultToken to environment variable VAULT_TOKEN
func _setVaultToken(token string) error {
	err := os.Setenv("VAULT_TOKEN", token)
	_checkError(err)
	return nil
}

//ReadKubeTokenFromFile and inject into env variable
func ReadKubeTokenFromFile(filename string) (string, error) {
	cb, err := ioutil.ReadFile(filename)
	_checkError(err)
	return string(cb), nil
}

// Login to Vault
func Login(address string, token string) error {
	type payload struct {
		Auth struct {
			ClientToken string `json:"client_token"`
		}
	}
	// form req
	var client http.Client
	buf := []byte(fmt.Sprintf(`{"jwt":"%s", "role": "kube-admin"}`, token))
	req, err := http.NewRequest("POST", address, bytes.NewBuffer(buf))
	_checkError(err)
	// post the req
	res, err := client.Do(req)
	_checkError(err)
	// format the response
	parse, err := ioutil.ReadAll(res.Body)
	_checkError(err)
	cb := payload{}
	err = json.Unmarshal(parse, &cb)
	_checkError(err)
	// set client token as env VAULT_TOKEN
	err = _setVaultToken(cb.Auth.ClientToken)
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
