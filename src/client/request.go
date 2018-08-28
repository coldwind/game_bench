package main

import (
	"net/http"
	"net/url"
	"sync/atomic"

	"github.com/bitly/go-simplejson"
)

func httpRequest(requestUrl, token, param string) {

	paramSet := make(map[string]string)

	atomic.AddInt32(&reqRes.Total, 1)

	if param != "" {
		json, err := simplejson.NewJson([]byte(param))
		if err == nil {
			jmap, err := json.Map()
			if err == nil {
				for k, v := range jmap {
					paramSet[k] = v.(string)
				}
			}
		}
	}

	postParam := url.Values{}
	if token != "" {
		postParam["token"] = append(postParam["token"], token)
	}
	for k, v := range paramSet {
		postParam[k] = append(postParam[k], v)
	}

	// 发送请求
	resp, err := http.PostForm(requestUrl, postParam)
	if err != nil {
		atomic.AddInt32(&reqRes.Failure, 1)
		atomic.AddInt32(&reqResByUrl[requestUrl].Failure, 1)
	} else {
		atomic.AddInt32(&reqRes.Success, 1)
		atomic.AddInt32(&reqResByUrl[requestUrl].Success, 1)
		if resp.StatusCode == 200 {
			atomic.AddInt32(&reqRes.HttpCode_200, 1)
			atomic.AddInt32(&reqResByUrl[requestUrl].HttpCode_200, 1)

		}

		if resp.ContentLength > 0 {
			atomic.AddInt64(&reqRes.TotalSize, resp.ContentLength)
			atomic.AddInt64(&reqResByUrl[requestUrl].TotalSize, resp.ContentLength)
		}
	}
}
