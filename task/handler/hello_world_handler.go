package handler

import (
	"core/logger"
	"time"
)

type HelloWorldHandler struct {
}

func (e *HelloWorldHandler) Execute() {
	logger.Info("hello world, time : %s", time.Now())
}
