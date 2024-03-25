package main

import (
	"os"

	"github.com/shramanpaul/golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}
