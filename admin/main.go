package main

import (
	_ "admin/router"
	"core/cmd"
)

func main() {
	cmd.Execute("liaz-admin")
}
