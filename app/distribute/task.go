/*
@Time : 2020/5/29 17:00
@Author : zxr
@File : task
@Software: GoLand
*/
package distribute

import (
	"fmt"
	"grabpixabay/config"
)

type Task struct {
	TaskName  string
	Host      string
	Type      string
	PicUrl    string
	HostUrl   string
	Page      int
	StartTime int64
}

func NewTask() *Task {
	return &Task{}
}

func RunTask(task *Task) {
	fmt.Println(config.GConf)
	fmt.Printf("%+v", task)

}
