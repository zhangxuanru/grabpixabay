/*
@Time : 2020/5/28 16:46
@Author : zxr
@File : main
@Software: GoLand
*/
package test

import (
	"testing"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

//logrus 测试
type AddHook struct {
	AppName string
}

func (a *AddHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (a *AddHook) Fire(entry *logrus.Entry) error {
	entry.Data["appNmae"] = a.AppName
	return nil
}

func TestLogRus(t *testing.T) {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        false,
		FieldsOrder:     []string{"id"},
		NoColors:        false,
		TimestampFormat: time.RFC3339,
	})
	hook := &AddHook{
		AppName: "this is test",
	}
	logrus.AddHook(hook)
	fields := logrus.WithFields(logrus.Fields{
		"name": "zxr",
		"id":   2,
	})
	fields.Info("aaww")
}
