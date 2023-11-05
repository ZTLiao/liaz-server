package main

import (
	"core/application"
	"core/cmd"
	"core/logger"
)

func init() {
	application.GetApp().SetName("liaz-admin")
}

func main() {
	cmd.Execute()
	logger.Info("hello,world!")
}
