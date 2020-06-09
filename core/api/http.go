/*
@Time : 2020/6/9 14:47
@Author : zxr
@File : http
@Software: GoLand
*/
package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const UA = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36"

//GET请求
func GetApi(url string, params *HttpParams) (bytes []byte, err error) {
	var (
		request  *http.Request
		response *http.Response
	)
	if request, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return
	}
	if params == nil {
		params = &HttpParams{
			Ua: UA,
		}
	}
	request.Header.Set("User-Agent", params.Ua)
	if response, err = http.DefaultClient.Do(request); err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("StatusCode=%d not ok", response.StatusCode))
	}
	defer response.Body.Close()
	bytes, err = ioutil.ReadAll(response.Body)
	return
}
