package main

import (
	"github.com/priteshbandi/notation-plugin-base/cli"
)

func main() {
	plg, _ := NewExamplePlugin()
	cli.NewHelper(plg).Execute()
}
