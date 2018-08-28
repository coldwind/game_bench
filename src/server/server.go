package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func serverStart() {
	http.HandleFunc("/test_start", testStart)
	http.HandleFunc("/test_stop", testStop)
	http.ListenAndServe(":19850", nil)
}

func testStart(res http.ResponseWriter, req *http.Request) {

	if recordFlag != 0 {
		res.Write([]byte("测试正在进行中"))
		return
	}

	recordFlag = 1

	// 打开日志记录文件
	t := time.Now().Unix()
	str := "bench:" + strconv.Itoa(int(t))
	sum := md5.Sum([]byte(str))
	filename := fmt.Sprintf("%x", sum)

	go record(filename)

	res.Write([]byte(filename))

}

func testStop(res http.ResponseWriter, req *http.Request) {

	recordFlag = 0

	res.Write([]byte("测试已停止"))

}

func record(filename string) {

	fp, err := os.OpenFile("/tmp/bench:"+filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 查看内容使用
	for {
		if recordFlag == 0 {
			break
		}

		date := time.Now().Format("2006-01-02 15:04:05")

		// 内存
		memInfo, err := mem.VirtualMemory()
		if err == nil {
			total := memInfo.Total / 1048576
			free := memInfo.Free / 1048576
			available := memInfo.Available / 1048576
			memString := fmt.Sprintf("[%s] [Memeory] Total: %v, Free:%v MB, Available:%v MB,  UsedPercent:%.2f %%\n", date, total, free, available, memInfo.UsedPercent)
			fp.WriteString(memString)
		}

		cpuInfo, err := cpu.Percent(time.Second, false)
		if err == nil {
			cpuString := fmt.Sprintf("[%s] [CPU] Used: %.2v %%\n", date, cpuInfo[0])
			fp.WriteString(cpuString)
		}

		// DISK
		diskInfo, err := disk.Usage("/")
		if err == nil {
			total := diskInfo.Total / 1048576
			free := diskInfo.Free / 1048576
			diskString := fmt.Sprintf("[%s] [Disk] Total: %.2v MB Free: %.2v MB\n", date, total, free)
			fp.WriteString(diskString)
		}

		// net
		netInfo, err := net.IOCounters(true)
		if err == nil {
			diskString := fmt.Sprintf("[%s] [Network] recv: %v bytes send: %v bytes\n", date, netInfo[0].BytesRecv, netInfo[0].BytesSent)
			fp.WriteString(diskString)
		}

		fp.WriteString("::::::::::::::::::::::::::::::\n")

		time.Sleep(2 * time.Second)
	}
}

func execCommand(command string) ([]byte, error) {

	cmd := exec.Command("/bin/bash", "-c", command)
	out, err := cmd.Output()

	if err != nil {
		return []byte{}, err
	}

	return out, nil
}
