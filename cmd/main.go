package main

import (
	"flag"
	"github.com/dimaglushkov/braingo/internal"
	"log"
)

func run() error {
	filePath := flag.String("r", "", "path to bf file to run")
	flag.Parse()

	if *filePath != "" {
		return internal.ExecuteFile(*filePath)
	}
	
	return internal.ExecuteInteractive()

}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
