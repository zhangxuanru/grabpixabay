/*
@Time : 2020/5/29 15:20
@Author : zxr
@File : flag
@Software: GoLand
*/
package initialize

import (
	"flag"
	"grabpixabay/configs"
	"time"

	"github.com/rifflock/lfshook"

	"github.com/pkg/errors"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

//命令行参数
type CommandLine struct {
	Type      string
	Order     string
	Color     string
	CountPage int
	Size      int
}

//命令行初始化
func Flag() *CommandLine {
	command := &CommandLine{}
	flag.StringVar(&command.Type, "type", configs.ImageType, "image:抓图片 video:抓视频")
	flag.StringVar(&command.Order, "order", "popular", "排序规则，latest:最新,popular:默认")
	flag.StringVar(&command.Color, "color", "", "按指定的颜色抓取，默认所有颜色"+
		"color values: grayscale, transparent, red, orange, yellow, green, turquoise, blue, lilac, pink, white, gray, black, brown")
	flag.IntVar(&command.CountPage, "count", 0, "请输入要抓取的总页数，默认是全部抓取")
	flag.IntVar(&command.Size, "size", 150, "请输入每页抓取的数量，默认100条")
	flag.Parse()
	return command
}

//初始化logrus
func Log() {
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
