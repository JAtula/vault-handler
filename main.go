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

package main

import (
	"log"
	"os"

	"github.com/urfave/cli.v2"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "token, t",
			Usage:    "token for vault auth",
			FilePath: "/var/run/secrets/kubernetes.io/serviceaccount/token",
			Destination: &token
		},
		cli.StringFlag{
			Name:  "path, p",
			Usage: "secret path `/secret/data/foo`",
			Destination: &path
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "output path `/opt/secret/secrets.json`",
			Destination: &output
		},

		app.Action = func(c *cli.Context) error {
			name := "someone"
			if c.NArg() > 0 {
			  name = c.Args()[0]
			}
			if language == "spanish" {
			  fmt.Println("Hola", name)
			} else {
			  fmt.Println("Hello", name)
			}
			return nil
		  }
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
