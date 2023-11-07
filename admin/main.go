package main

import (
	_ "admin/router"
	"core/cmd"
	"core/logger"
)

func main() {
	cmd.Execute()
	logger.Info("hello,world!")
}
