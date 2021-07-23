package main

import (
	"github.com/GeonHyeok-Lee/nomadcoin/cli"
	"github.com/GeonHyeok-Lee/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
