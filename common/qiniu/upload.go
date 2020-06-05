/*
@Time : 2020/6/4 18:56
@Author : zxr
@File : file
@Software: GoLand
*/
package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"grabpixabay/config"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

//上传文件
//var (
//	accessKey = os.Getenv("QINIU_ACCESS_KEY")
//	secretKey = os.Getenv("QINIU_SECRET_KEY")
//	bucket    = os.Getenv("QINIU_TEST_BUCKET")
//)

type QiNiu struct {
	UA         string
	Referer    string
	SrcFile    string
	UpFileName string
}

func NewQiNiu() *QiNiu {
	return &QiNiu{}
}

//上传文件
func (q *QiNiu) UploadHttpFile() (putRet *storage.PutRet, err error) {
	putPolicy := storage.PutPolicy{
		Scope: config.QINIU_BKT + ":" + q.UpFileName,
	}
	mac := qbox.NewMac(config.QINIU_AK, config.QINIU_SK)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	putRet = &storage.PutRet{}
	putExtra := storage.PutExtra{}
	if strings.Contains(q.SrcFile, "http") {
		var data []byte
		if data, err = q.DownloadFile(); err != nil {
			return
		}
		dataLen := int64(len(data))
		err = formUploader.Put(context.Background(), putRet, upToken, q.UpFileName, bytes.NewReader(data), dataLen, &putExtra)
	} else {
		err = formUploader.PutFile(context.Background(), putRet, upToken, q.UpFileName, q.SrcFile, &putExtra)
	}
	if err != nil {
		return
	}
	return putRet, nil
}

//从URL中取文件名
func (q *QiNiu) GenDefaultFileName() string {
	fileName := path.Base(q.SrcFile)
	return fileName
}

//下载文件
func (q *QiNiu) DownloadFile() (bytes []byte, err error) {
	var (
		request  *http.Request
		response *http.Response
	)
	if request, err = http.NewRequest(http.MethodGet, q.SrcFile, nil); err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", q.UA)
	request.Header.Set("referer", q.Referer)
	if response, err = http.DefaultClient.Do(request); err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("http code is:%d", response.StatusCode))
	}
	bytes, err = ioutil.ReadAll(response.Body)
	return
}
