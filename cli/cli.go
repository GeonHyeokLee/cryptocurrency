package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/GeonHyeok-Lee/cryptocurrency/explorer"
	"github.com/GeonHyeok-Lee/cryptocurrency/rest"
)

func usage() {
	fmt.Printf("\n\n-------------------------------------------------------------\n\n")
	fmt.Printf("Welcome to CryptoCurrency!\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("-port: 			the PORT of the server\n")
	fmt.Printf("-mode: 			Choose between 'html' and 'rest'\n\n")
	fmt.Printf("-------------------------------------------------------------\n\n")
	// runtime.Goexit()
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
