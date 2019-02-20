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

	"github.com/JAtula/vault-handler/pkg/secrets"
	"github.com/JAtula/vault-handler/pkg/util"
	"gopkg.in/urfave/cli.v1"
)

func _errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

// Execute main app
func Execute() {
	app := cli.NewApp()
	app.Name = "vault-handler"
	app.Description = "Read key-value secrets from Hashicorp's Vault."
	app.Usage = "Yet another Vault CLI tool."
	app.Version = "0.2.0"
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

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command %q not found.\n\n", command)
		cli.ShowAppHelpAndExit(c, 1)
	}

	app.Commands = []cli.Command{
		{
			Name:    "read",
			Aliases: []string{"r"},
			Usage:   "read secret path `/secret/data/test`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "token-path, t",
					Usage:       "(Required) token path `./auth_token`",
					Destination: &token,
				},
				cli.StringFlag{
					Name:        "path, p",
					Usage:       "(Required) secret path `/secret/data/test`",
					Destination: &path,
				},
				cli.StringFlag{
					Name:        "output, o",
					Usage:       "output path `.env`",
					Destination: &output,
				},
			},
			CustomHelpTemplate: `
OPTIONS:
{{range .VisibleFlags}}{{ "\t" }}{{.}}{{ "\n" }}{{end}}
			`,
			SkipFlagParsing: false,
			HideHelp:        false,
			Hidden:          false,
			Action: func(c *cli.Context) error {
				if token == "" || path == "" {
					fmt.Fprintf(c.App.Writer, "\nRequired flags missing.\n")
					cli.ShowCommandHelpAndExit(c, "read", 1)
				}
				tokenData, err := util.ReadTokenFromFile(token)
				_errCheck(err)
				err = util.SetVaultToken(tokenData)
				_errCheck(err)
				data, err := secrets.ReadSecret(path, string(os.Getenv("VAULT_ADDR")))
				_errCheck(err)
				if data == "404" {
					return cli.NewExitError("Secret not found.", 1)
				}
				if output != "" {
					err = util.WriteSecretPayload(output, data)
					_errCheck(err)
					return nil
				}
				fmt.Println(data)
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
