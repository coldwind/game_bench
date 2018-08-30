package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sync"
	"time"
)

// 请求结果记录
type RequestResult struct {
	Failure      int32
	Success      int32
	HttpCode_200 int32
	TotalSize    int64
	Total        int32
}

// token数据
var tokenData []string
var reqRes *RequestResult
var reqResByUrl map[string]*RequestResult

func main() {
	client := flag.Int("c", 100, "客户端并发数量")
	total := flag.Int("n", 1000, "请求总量")
	//request := flag.String("r", "http", "请求类型: http(s)/ws")
	address := flag.String("a", "", "请求地址: http(http://www.domain.com) | ws(ws://domain.com)")
	useToken := flag.String("t", "", "使用token文件[yes|no]")
	serverAddress := flag.String("s", "", "服务器端监控地址(http://www.domain.com:19850)")

	flag.Parse()

	if *client > *total {
		log.Println("并发数量不能大于总请求量")
		return
	}

	reqRes = new(RequestResult)
	reqResByUrl = make(map[string]*RequestResult)

	// 根据类型启动压测协程
	totalF := float64(*total)
	clientF := float64(*client)
	groupNum := int(math.Ceil(totalF / clientF))

	var tokenLen int
	if *useToken == "yes" {
		tokenData, _ = tokenList()
		tokenLen = len(tokenData)
	} else {
		tokenData = make([]string, 0, 1)
		tokenLen = 0
	}

	// 分配请求链接
	var requestLen int
	var rData *RequestStruct

	if *address == "" {

		rData, _ = readReqeustYaml()
		requestLen = len(rData.Request)

		for _, v := range rData.Request {
			if _, ok := reqResByUrl[v.Address]; !ok {
				reqResByUrl[v.Address] = new(RequestResult)
			}
		}
	} else {
		rData = new(RequestStruct)
		rData.Request = make([]*RequestParamStruct, 1)
		rData.Request[0] = new(RequestParamStruct)
		rData.Request[0].Address = *address
		rData.Request[0].Param = "{}"
		requestLen = 1
		reqResByUrl[*address] = new(RequestResult)
	}

	wg := new(sync.WaitGroup)

	// 请求服务器端监控状态
	if *serverAddress != "" {
		resp, err := http.Get(*serverAddress + "/test_start")
		if err == nil {
			serverBody, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				OutputLine("SESSION KEY")
				fmt.Println(string(serverBody))
			}
		}
	}

	OutputLine("请求概况")

	totalBegin := time.Now().UnixNano() / 1e6
	for j := 1; j <= groupNum; j++ {
		wg.Add(int(clientF))

		// 获取执行开始时间
		begin := time.Now().UnixNano() / 1e6
		var token string
		for i := 0; i < *client; i++ {
			go func(g, c int) {
				// 循环
				gc := g * c
				if tokenLen > 0 {
					tkey := gc % tokenLen
					token = tokenData[tkey]
				} else {
					token = ""
				}
				key := gc % requestLen
				httpRequest(rData.Request[key].Address, token, rData.Request[key].Param)

				wg.Done()
			}(j, i)
		}

		wg.Wait()

		// 获取执行结束时间
		end := time.Now().UnixNano() / 1e6
		OutputPerGroup(begin, end, clientF, j)
	}

	// 请求服务器端监控状态
	if *serverAddress != "" {
		resp, err := http.Get(*serverAddress + "/test_stop")
		if err == nil {
			serverBody, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				OutputLine("SESSION KEY")
				fmt.Println(string(serverBody))
			}
		}
	}

	totalEnd := time.Now().UnixNano() / 1e6
	OutputTotal(totalBegin, totalEnd)

	OutputLine("接口详情")
	OutputUrlInfo()
}
