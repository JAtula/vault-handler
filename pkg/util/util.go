package util

import "github.com/hashicorp/vault/api"

func login(token string, addr string) (string, error) {
	conf := &api.Config{
		Address: addr,
	}
	client, err := api.NewClient(conf)
	if err != nil {
		return "", err
	}
	client.SetToken(token)
	return token, nil
}
