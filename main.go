package main

import (
	"bufio"
	_ "embed"
	"github.com/dimaglushkov/braingo/internal"
	"io"
	"log"
	"os"
)

//go:embed example/hello_world.bf
var code string

func run() error {
	//	interactive mode
	in := bufio.NewReader(os.Stdin)
	var x string
	var err error
	for {
		if x, err = in.ReadString('\n'); err != nil {
			if err == io.EOF {
				return nil
			}
		}
		if x[0] == '\\' {
			if err = internal.ExecInteractive(x); err != nil {
				return err
			}
		} else {
			if err = internal.Execute(x); err != nil {
				return err
			}
		}
	}

}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
