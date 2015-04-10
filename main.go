package main

import (
	"github.com/medimatrix/stacky/stacky"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("You must specify a stacky file to run.")
	}

	f := os.Args[1]
	vm := new(stacky.VM)
	instructions, err := stacky.Read(f)
	if err != nil {
		log.Fatalln(err)
	}

	vm.Interpret(instructions)
}
