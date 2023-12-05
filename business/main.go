package main

import (
	_ "business/router"
	"core/cmd"
)

func main() {
	cmd.Execute("liaz-business")
}
