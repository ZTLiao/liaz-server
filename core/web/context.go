package web

import (
	"context"
	"core/application"
	"core/constant"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebContext struct {
	Context *gin.Context
}

func (e *WebContext) Trace(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.WithFields(logrus.Fields{
		constant.REQUEST_ID: e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Tracef(format, args...)
}

func (e *WebContext) Debug(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.WithFields(logrus.Fields{
		constant.REQUEST_ID: e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Debugf(format, args...)
}

func (e *WebContext) Info(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.WithFields(logrus.Fields{
		constant.REQUEST_ID: e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Infof(format, args...)
}

func (e *WebContext) Warn(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.WithFields(logrus.Fields{
		constant.REQUEST_ID: e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Warnf(format, args...)
}

func (e *WebContext) Error(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.WithFields(logrus.Fields{
		constant.REQUEST_ID: e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Errorf(format, args...)
}

func (ctx *WebContext) Fatal(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.WithFields(logrus.Fields{
		constant.REQUEST_ID: ctx.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Fatalf(format, args...)
}

func (e *WebContext) Panic(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.WithFields(logrus.Fields{
		constant.REQUEST_ID: e.Context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Panicf(format, args...)
}

func (e *WebContext) Background() context.Context {
	return e.Context.Request.Context()
}
