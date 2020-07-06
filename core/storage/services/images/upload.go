package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"grabpixabay/configs"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

type QiNiu struct {
	UA         string
	Referer    string
	SrcFile    string
	UpFileName string
}

type UploadResult struct {
	Size     int64
	FileName string
	PutRet   *storage.PutRet
}

const DefaultUa = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36"

func NewQiNiu() *QiNiu {
	return &QiNiu{}
}

//上传文件
func (q *QiNiu) UploadFile() (ret *UploadResult, err error) {
	ret = &UploadResult{}
	if q.UpFileName == "" {
		q.UpFileName = q.GenDefaultFileName()
		ret.FileName = q.UpFileName
	}
	if q.SrcFile == "" {
		return ret, errors.New("srcFile is nil")
	}
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
	cfg.UseCdnDomains = true
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	putRet := &storage.PutRet{}
	putExtra := storage.PutExtra{}

	if strings.Contains(q.SrcFile, "http") {
		var data []byte
		if data, err = q.DownloadFile(); err != nil {
			return
		}
		dataLen := int64(len(data))
		ret.Size = dataLen
		err = formUploader.Put(context.Background(), putRet, upToken, q.UpFileName, bytes.NewReader(data), dataLen, &putExtra)
	} else {
		err = formUploader.PutFile(context.Background(), putRet, upToken, q.UpFileName, q.SrcFile, &putExtra)
	}
	if err != nil {
		return
	}
	ret.PutRet = putRet
	return ret, nil
}

//从URL中取文件名
func (q *QiNiu) GenDefaultFileName() string {
	if strings.Contains(q.SrcFile, "http") {
		var buff bytes.Buffer
		r, _ := url.Parse(q.SrcFile)
		filePath := strings.ReplaceAll(r.Path, "/", "")
		pathLen := len(filePath)
		if pathLen < 25 {
			return filePath
		}
		randSource := "0123456789abcdefghijklmnopqrstuvwxyz"
		sourceLen := len(randSource) - 1
		rand.Seed(time.Now().UnixNano())
		n1 := rand.Intn(sourceLen)
		n2 := rand.Intn(sourceLen)
		buff.WriteByte(randSource[n1])
		buff.WriteByte(randSource[n2])
		fileName := buff.String() + filePath[pathLen-25:]
		return fileName
	}
	fileName := path.Base(q.SrcFile)
	return fileName
}

//下载文件
func (q *QiNiu) DownloadFile() (bytes []byte, err error) {
	var (
		request  *http.Request
		response *http.Response
	)
	//如果不存在就不请求具体内容了
	//if exists, err := q.HeadExists(); err != nil || exists == false {
	//	return nil, err
	//}
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
	if response, err = q.NewClient().Do(request); err != nil {
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
	if response, err = q.NewClient().Do(request); err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return false, errors.New(fmt.Sprintf("http code is:%d", response.StatusCode))
	}
	return true, nil
}

func (q *QiNiu) NewClient() *http.Client {
	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	return c
}
