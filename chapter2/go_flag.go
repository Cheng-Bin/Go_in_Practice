package chapter2

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short: "n" long: "name"    default: "World"    description: "A name to Say hello to."`
	Spanish bool   `short: "s"    long: "spanish"   description: "use Spanish Language."`
}

func FlagParse() {
	flags.Parse(&opts)

	if opts.Spanish {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s !\n", opts.Name)
	}
}
