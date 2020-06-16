package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"grabpixabay/configs"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

type QiNiu struct {
	UA         string
	Referer    string
	SrcFile    string
	UpFileName string
}

const DefaultUa = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36"

func NewQiNiu() *QiNiu {
	return &QiNiu{}
}

//上传文件
func (q *QiNiu) UploadFile() (putRet *storage.PutRet, err error) {
	putPolicy := storage.PutPolicy{
		Scope: configs.QINIU_BKT + ":" + q.UpFileName,
	}
	mac := qbox.NewMac(configs.QINIU_AK, configs.QINIU_SK)
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
	if q.UpFileName == "" {
		q.UpFileName = q.GenDefaultFileName()
	}
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
	if q.UA == "" {
		q.UA = DefaultUa
	}
	if q.Referer != "" {
		request.Header.Set("referer", q.Referer)
	}
	request.Header.Set("User-Agent", q.UA)
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

//先head请求 判断文件是否存在
func (q *QiNiu) HeadExists() (exists bool, err error) {
	var (
		request  *http.Request
		response *http.Response
	)
	if request, err = http.NewRequest(http.MethodHead, q.SrcFile, nil); err != nil {
		return false, err
	}
	if q.UA == "" {
		q.UA = DefaultUa
	}
	if q.Referer != "" {
		request.Header.Set("referer", q.Referer)
	}
	request.Header.Set("User-Agent", q.UA)
	if response, err = http.DefaultClient.Do(request); err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return false, errors.New(fmt.Sprintf("http code is:%d", response.StatusCode))
	}
	return true, nil
}
