package config

import (
	"core/system"
	"core/utils"
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Type   string `yaml:"type"`
	Path   string `yaml:"path"`
	Level  string `yaml:"level"`
	Stdout string `yaml:"stdout"`
}

func (e *Logger) Init() {
	log := logrus.New()
	if e != nil {
		switch e.Type {
		case "json":
			log.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat: utils.NORM_DATETIME_MS_PATTERN,
			})
		case "text":
			log.SetFormatter(&logrus.TextFormatter{
				TimestampFormat: utils.NORM_DATETIME_MS_PATTERN,
			})
		default:
			log.SetFormatter(&logrus.TextFormatter{
				TimestampFormat: utils.NORM_DATETIME_MS_PATTERN,
			})
		}
		if e.Stdout == "file" {
			writer, err := rotatelogs.New(
				e.Path+".%Y%m%d%H%M",
				rotatelogs.WithLinkName(e.Path),
				//保留一周
				rotatelogs.WithMaxAge(time.Duration(7*24)*time.Hour),
				//30分钟分割一次
				rotatelogs.WithRotationTime(time.Duration(30)*time.Minute),
				rotatelogs.WithRotationSize(int64(1*1024*35000*1024)),
			)
			if err == nil {
				log.SetOutput(writer)
			} else {
				fmt.Println(err.Error())
				return
			}
		} else {
			log.SetOutput(os.Stdout)
		}
		switch e.Level {
		case "trace":
			log.SetLevel(logrus.TraceLevel)
		case "debug":
			log.SetLevel(logrus.DebugLevel)
		case "info":
			log.SetLevel(logrus.InfoLevel)
		case "warn":
			log.SetLevel(logrus.WarnLevel)
		case "error":
			log.SetLevel(logrus.ErrorLevel)
		case "fatal":
			log.SetLevel(logrus.FatalLevel)
		}
	}
	system.SetLogger(log)
}
