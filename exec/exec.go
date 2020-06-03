/*
@Time : 2020/5/29 16:00
@Author : zxr
@File : exec
@Software: GoLand
*/
package exec

import (
	"context"
	"grabpixabay/common/verify"

	"github.com/sirupsen/logrus"
)

func init() {
	InitLog()
}

func Run() {
	task := InitFlag()
	if err := verify.CheckTask(task); err != nil {
		logrus.WithFields(logrus.Fields{
			"Host": task.Host,
			"Type": task.Type,
		}).Error(err)
		return
	}
	task.Cxt, task.Can = context.WithCancel(context.Background())
	NotifySing(task.Can)
	task.RunTask()
}
