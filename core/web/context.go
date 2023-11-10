package web

import (
	"core/application"
	"core/constant"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebContext struct {
	Context *gin.Context
}

func (e *WebContext) Trace(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Tracef(format, args...)
}

func (e *WebContext) Debug(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Debugf(format, args...)
}

func (e *WebContext) Info(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Infof(format, args...)
}

func (e *WebContext) Warn(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Warnf(format, args...)
}

func (e *WebContext) Error(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Errorf(format, args...)
}

func (ctx *WebContext) Fatal(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": ctx.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Fatalf(format, args...)
}

func (e *WebContext) Panic(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.WithFields(logrus.Fields{
		"requestId": e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Panicf(format, args...)
}
