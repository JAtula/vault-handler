package cmd

import (
	"flag"
	"fmt"
)

// Parse input arguments
func Parse() {
	secretFlag := flag.String("secret-path", "", "vault-handler -secret-path /secret/data/test")
	pathFlag := flag.String("output", "", "vault-handler -secret-path /secret/data/test -output ./secret-data.json")
	boolFlag := flag.Bool("token", false, "vault-handler -secret-path /secret/data/test -output ./secret-data.json -token supersecrettoken")

	flag.Parse()
	fmt.Println("secretFlag: ", *secretFlag)
	fmt.Println("pathFlag: ", *pathFlag)
	fmt.Println("boolFlag: ", *boolFlag)
}
