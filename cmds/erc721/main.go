package main

import (
	"fmt"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc721/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
