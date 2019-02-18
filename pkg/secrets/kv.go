package secrets

import (
	"encoding/json"

	"github.com/hashicorp/vault/api"
)

func parseJSON(parse map[string]interface{}) (string, error) {
	pd, err := json.Marshal(parse)
	if err != nil {
		return "", err
	}
	return string(pd), nil
}

// ReadSecret from Hashicorp Vault
func ReadSecret(path string, addr string) (string, error) {
	conf := &api.Config{
		Address: addr,
	}
	c, err := api.NewClient(conf)
	if err != nil {
		return "", err
	}
	client := *c.Logical()
	cb, err := client.Read(path)
	if err != nil {
		return "", err
	}
	if cb == nil {
		return "404", nil
	}
	parsed, err := parseJSON(cb.Data)
	if err != nil {
		return "", err
	}
	return parsed, nil
}
