package main

import (
	"github.com/notaryproject/notation-plugin-base/cli"
	"github.com/notaryproject/notation-plugin-base/dummyplugin"
)

func main() {
	plg, _ := sampleplugin.NewDummyPlugin()
	cli.NewHelper(plg).Execute()
}
