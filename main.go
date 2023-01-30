package main

import (
	_ "embed"
	"github.com/dimaglushkov/braingo/internal"
	"log"
)

//go:embed example/hello_world.bf
var code string

func run() error {
	return internal.Execute(code)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
