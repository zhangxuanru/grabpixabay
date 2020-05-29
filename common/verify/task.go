/*
@Time : 2020/5/29 18:00
@Author : zxr
@File : task
@Software: GoLand
*/
package verify

import (
	"errors"
	"fmt"
	"grabpixabay/app/distribute"
	"grabpixabay/common/util"
	"grabpixabay/config"
)

//验证task数据完整性
func CheckTask(task *distribute.Task) (err error) {
	GConf := config.GConf
	task.HostUrl = GConf.HostMap[task.Host]
	if task.HostUrl == "" {
		return errors.New("host 错误,目前仅支持" + config.PIX_HOST)
	}
	if _, ok := GConf.TypeMap[task.Type]; !ok {
		errStr := fmt.Sprintf("type 错误,目前仅支持:%s,%s,%s,%s", config.TYPE_ALL, config.TYPE_LATEST, config.TYPE_SIFT, config.TYPE_PIC)
		return errors.New(errStr)
	}
	if task.Page < 1 {
		task.Page = 0
	}
	task.StartTime = util.GetNowUnixTime()
	return nil
}
