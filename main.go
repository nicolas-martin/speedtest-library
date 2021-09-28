package main

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

func main() {
	user, _ := speedtest.FetchUserInfo()
	fmt.Println("== users ==")
	fmt.Println(user)

	serverList, _ := speedtest.FetchServerList(user)
	fmt.Println("== server list ==")
	fmt.Println(serverList)

	targets, _ := serverList.FindServer([]int{})
	fmt.Println("== targets ==")
	fmt.Println(targets)

	for i, s := range targets {
		fmt.Println(" ** ** ** ** ")
		fmt.Printf("running %d %s\n", i, s)
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)

		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
	}
}
