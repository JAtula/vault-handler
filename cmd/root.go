// Copyright © 2018 Juhani Atula juhaniatula@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/JAtula/vault-handler/pkg/util"

	"github.com/JAtula/vault-handler/pkg/secrets"

	"gopkg.in/urfave/cli.v1"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Execute main app
func Execute() {
	app := cli.NewApp()
	app.Name = "vault-handler - handle key-value secrets in Hashicorp's Vault"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Juhani Atula",
			Email: "juhani.atula@polarsquad.com",
		},
	}

	var (
		token  string
		path   string
		output string
	)

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "token, t",
			Usage:       "token path",
			Value:       "/var/run/secrets/kubernetes.io/serviceaccount/token",
			Destination: &token,
		},
		cli.StringFlag{
			Name:        "path, p",
			Usage:       "secret path `/secret/data/foo.json`",
			Destination: &path,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "output path `/opt/secret/secrets.json`",
			Destination: &output,
		},
	}

	app.Action = func(c *cli.Context) error {
		tokenData, err := util.ReadTokenFromFile(token)
		check(err)
		err = util.SetVaultToken(tokenData)
		check(err)
		data, err := secrets.ReadSecret(path, string(os.Getenv("VAULT_ADDR")))
		check(err)
		if data == "404" {
			fmt.Println("Secret not found.")
			os.Exit(1)
		}
		cb, err := util.WriteSecretPayload(output, data)
		check(err)
		fmt.Println(cb)
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
