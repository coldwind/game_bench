package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func OutputPerGroup(begin, end int64, client float64, index int) {
	clientInt := int(client)
	offset := (index - 1) * clientInt
	log.Printf("第 %d-%d 次请求, 耗时: %d ms", offset, clientInt+offset, end-begin)
}

func OutputTotal(begin, end int64) {
	useTime := end - begin
	rate := math.Floor(float64(reqRes.Total) / float64(useTime) * 1000)
	fmt.Printf("总请求:%d 总耗时:%d ms 吞吐率:%d/s", reqRes.Total, useTime, int64(rate))
}

func OutputLine(title string) {
	fmt.Println("")
	fmt.Println(strings.Repeat("-", 30), title, strings.Repeat("-", 30))
}

func OutputUrlInfo() {
	for k, v := range reqResByUrl {
		fmt.Printf("URL:%s, 数据量:%d 成功:%d 失败:%d 200状态:%d\n", k, v.TotalSize, v.Success, v.Failure, v.HttpCode_200)
	}
}
