/*
@Time : 2020/5/29 16:00
@Author : zxr
@File : exec
@Software: GoLand
*/
package exec

import (
	"fmt"
	"grabpixabay/app/distribute"
	"grabpixabay/common/verify"
	"grabpixabay/config"

	"github.com/sirupsen/logrus"
)

func Init() {
	InitFlag()
	InitLog()
}

func Run() {
	fmt.Printf("%v\n\n", config.FULL_NAME)
	Init()
	config.AppConfig()
	if err := verify.CheckTask(Task); err != nil {
		logrus.WithFields(logrus.Fields{
			"Host": Task.Host,
			"Type": Task.Type,
		}).Error(err)
		return
	}
	distribute.RunTask(Task)
}
