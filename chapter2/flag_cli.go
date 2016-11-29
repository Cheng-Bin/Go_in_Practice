package chapter2

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name yo say hello to.")
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use spanish language")
	flag.BoolVar(&spanish, "s", false, "Use spanish language.")
}

func Parse() {

	flag.VisitAll(func(flag *flag.Flag) {

		format := "\t-%s: %s (Default: '%s')\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)

	})

	flag.Parse()

	if spanish {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
