package main

import (
	"github.com/GeonHyeok-Lee/cryptocurrency/cli"
	"github.com/GeonHyeok-Lee/cryptocurrency/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
