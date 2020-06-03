/*
@Time : 2020/5/29 15:20
@Author : zxr
@File : flag
@Software: GoLand
*/
package exec

import (
	"context"
	"flag"
	"fmt"
	"grabpixabay/app/distribute"
	"grabpixabay/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rifflock/lfshook"

	"github.com/pkg/errors"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

//命令行初始化
func InitFlag() *distribute.Task {
	task := distribute.NewTask()
	flag.StringVar(&task.Host, "host", config.PIX_HOST, "请输入要抓取的host,目前仅支持pixabay")
	flag.IntVar(&task.Page, "page", 0, "请输入要抓取的页数，默认是全部抓取")
	flag.StringVar(&task.Type, "type", config.TYPE_ALL, "all:全站抓取 latest:获取最新,sift:获取小编精选,pic:获取图片详情")
	flag.StringVar(&task.PicUrl, "pic", "", "图片详情页地址")
	flag.Parse()
	return task
}

//初始化logrus
func InitLog() {
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: time.RFC3339,
	})
	logrus.SetLevel(logrus.DebugLevel)
	//logrus.SetReportCaller(true)

	//日志切割
	path := "./logs/error.log"
	writer, err := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path),             // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(5*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	pathInfo := "./logs/info.log"
	writerInfo, _ := rotatelogs.New(
		pathInfo+".%Y%m%d",
		rotatelogs.WithLinkName(pathInfo),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(5*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writerInfo, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writerInfo,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, nil)
	logrus.AddHook(lfHook)
}

//监听信号
func NotifySing(cancel context.CancelFunc) {
	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
	go func() {
		sig := <-sign
		fmt.Println("接收到信号:", sig)
		cancel()
	}()
}
