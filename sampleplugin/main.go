package main

import (
	"github.com/priteshbandi/notation-plugin-base/cli"
	"github.com/priteshbandi/notation-plugin-base/sampleplugin"
)

func main() {
	plg, _ := sampleplugin.NewSamplePlugin()
	cli.NewHelper(plg).Execute()
}
