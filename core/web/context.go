package web

import (
	"context"
	"core/constant"
	"core/system"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebContext struct {
	context *gin.Context
}

func NewWebContext(context *gin.Context) *WebContext {
	return &WebContext{context}
}

func (e *WebContext) Trace(format string, args ...interface{}) {
	system.GetLogger().WithFields(logrus.Fields{
		constant.REQUEST_ID: e.context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Tracef(format, args...)
}

func (e *WebContext) Debug(format string, args ...interface{}) {
	system.GetLogger().WithFields(logrus.Fields{
		constant.REQUEST_ID: e.context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Debugf(format, args...)
}

func (e *WebContext) Info(format string, args ...interface{}) {
	system.GetLogger().WithFields(logrus.Fields{
		constant.REQUEST_ID: e.context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Infof(format, args...)
}

func (e *WebContext) Warn(format string, args ...interface{}) {
	system.GetLogger().WithFields(logrus.Fields{
		constant.REQUEST_ID: e.context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Warnf(format, args...)
}

func (e *WebContext) Error(format string, args ...interface{}) {
	system.GetLogger().WithFields(logrus.Fields{
		constant.REQUEST_ID: e.context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Errorf(format, args...)
}

func (ctx *WebContext) Fatal(format string, args ...interface{}) {
	system.GetLogger().WithFields(logrus.Fields{
		constant.REQUEST_ID: ctx.context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Fatalf(format, args...)
}

func (e *WebContext) Panic(format string, args ...interface{}) {
	system.GetLogger().WithFields(logrus.Fields{
		constant.REQUEST_ID: e.context.Request.Header.Get(constant.X_REQUEST_ID),
	}).Panicf(format, args...)
}

func (e *WebContext) Background() context.Context {
	return e.context.Request.Context()
}

func (e *WebContext) AbortWithError(err error) {
	e.Error(err.Error())
	e.context.AbortWithError(http.StatusInternalServerError, err)
}

func (e *WebContext) ClientIP() string {
	return e.context.ClientIP()
}

func (e *WebContext) GetHeader(key string) string {
	return e.context.Request.Header.Get(key)
}

func (e *WebContext) PostForm(key string) string {
	return e.context.PostForm(key)
}

func (e *WebContext) Param(key string) string {
	return e.context.Param(key)
}

func (e *WebContext) Query(key string) string {
	return e.context.Query(key)
}

func (e *WebContext) DefaultQuery(key string, defaultValue string) string {
	return e.context.DefaultQuery(key, defaultValue)
}

func (e *WebContext) ShouldBindJSON(obj any) error {
	return e.context.ShouldBindJSON(obj)
}

func (e *WebContext) BindJSON(obj any) error {
	return e.context.BindJSON(obj)
}

func (e *WebContext) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return e.context.Request.FormFile(key)
}

func (e *WebContext) MultipartForm() (*multipart.Form, error) {
	return e.context.MultipartForm()
}

func GetUserId(wc *WebContext) int64 {
	userIdStr := wc.GetHeader(constant.X_USER_ID)
	if len(userIdStr) == 0 {
		return 0
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		wc.Error(err.Error())
		return 0
	}
	return userId
}
