package main

import (
	"os"

	"github.com/jaehong-hwang/todo/cli"
)

func main() {
	app := cli.NewApp()

	app.Run(os.Args).Print()
}
