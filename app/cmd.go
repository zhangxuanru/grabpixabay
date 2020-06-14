/*
@Time : 2020/6/8 15:59
@Author : zxr
@File : cmd
@Software: GoLand
*/
package app

import (
	"context"
	"errors"
	"grabpixabay/configs"
	"grabpixabay/core/scheduler"
	"grabpixabay/initialize"
	"grabpixabay/util"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

func init() {
	initialize.Log()
	configs.AppConfig()
}

var (
	once       sync.Once
	command    *initialize.CommandLine
	concurrent *scheduler.Concurrent
	err        error
)

//程序入口
func Run() {
	if command, err = verifyCommand(); err != nil {
		logrus.Error(err)
		return
	}
	task := buildTask(command)
	once.Do(func() {
		concurrent = scheduler.NewConcurrent(configs.GConf.WorkerCount, task.Ctx, task.Can)
		concurrent.Run()
		task.Pool = concurrent
	})
	if command.Type == configs.ImageType {
		task.CallImage()
	} else {
		task.CallVideo()
	}
	defer concurrent.Wait()
	//item.Monitor()
}

//构建task
func buildTask(command *initialize.CommandLine) *scheduler.Task {
	ctx, cancel := context.WithCancel(context.Background())
	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
	task := &scheduler.Task{
		Command:  command,
		Ctx:      ctx,
		Can:      cancel,
		SignChan: sign,
	}
	return task
}

//验证命令行参数
func verifyCommand() (command *initialize.CommandLine, err error) {
	command = initialize.Flag()
	if command.Type != configs.ImageType && command.Type != configs.VideoType {
		return command, errors.New("type undefined")
	}
	if command.Color != "" && util.InStrings(configs.GConf.Colors, command.Color) == false {
		return command, errors.New("color undefined")
	}
	if command.CountPage < 0 {
		return command, errors.New("CountPage It can't be less than 0")
	}
	if command.Size > 200 {
		return command, errors.New("Size Accepted values: 3 - 200 ")
	}
	return command, nil
}
