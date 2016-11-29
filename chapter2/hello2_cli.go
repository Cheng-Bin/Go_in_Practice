package chapter2

import (
	cli "gopkg.in/urfave/cli.v1"
)

func Hello2Cli() {
	app := cli.NewApp()
	app.Usage = "Count up or down"
}
