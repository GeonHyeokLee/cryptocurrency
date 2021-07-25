package main

import (
	"github.com/GeonHyeok-Lee/minimal-cryptocurrency/cli"
	"github.com/GeonHyeok-Lee/minimal-cryptocurrency/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
