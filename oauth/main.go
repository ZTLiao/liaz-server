package main

import (
	"core/cmd"
	_ "oauth/router"
)

func main() {
	cmd.Execute("liaz-oauth")
}
