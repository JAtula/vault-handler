# Vault Handler
Vault handler is dockerized handler to be used as an init container with a parent container that it feeds secrets to. 

The CLI tool only reads a key-value secret and writes the output to a JSON file. 

## Build

Download the latest [release](https://github.com/JAtula/vault-handler/releases) and build it `go build && chmod +x ./vault-handler && mv ./vault-handler /usr/local/bin/`. 

## Usage

```
GLOBAL OPTIONS:
   --output /opt/secret/secrets.json, -o /opt/secret/secrets.json  output path /opt/secret/secrets.json
   --path /secret/data/foo.json, -p /secret/data/foo.json          secret path /secret/data/foo.json
   --token value, -t value                                         token path (default: "/var/run/secrets/kubernetes.io/serviceaccount/token")
   --help, -h                                                      show help
   --version, -v                                                   print the version
```

This is in early development and shouldn't be used in production.