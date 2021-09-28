package main

import (
	"fmt"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

func main() {
	user, _ := speedtest.FetchUserInfo()
	fmt.Println("== users ==")
	fmt.Println(user)
	fmt.Println("== users ==")

	serverList, _ := speedtest.FetchServerList(user)
	fmt.Println("== server list ==")
	fmt.Println(serverList)
	fmt.Println("== server list ==")

	targets, _ := serverList.FindServer([]int{})
	fmt.Println("== targets ==")
	fmt.Println(targets)
	fmt.Println("== targets ==")

	for i, s := range targets {
		fmt.Println("== run ==")
		fmt.Printf("running %d %s %s\n", i, s, s.Country)
		fmt.Println("== run ==")
		s.PingTest()
		testDownload(s, false)
		testUpload(s, false)

		fmt.Printf("\nLatency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
	}
}
func testDownload(server *speedtest.Server, savingMode bool) error {
	quit := make(chan bool)
	fmt.Printf("Download Test: ")
	go dots(quit)
	err := server.DownloadTest(savingMode)
	quit <- true
	if err != nil {
		return err
	}
	fmt.Println()
	return err
}

func testUpload(server *speedtest.Server, savingMode bool) error {
	quit := make(chan bool)
	fmt.Printf("Upload Test: ")
	go dots(quit)
	err := server.UploadTest(savingMode)
	quit <- true
	if err != nil {
		return err
	}
	fmt.Println()
	return nil
}

func dots(quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			time.Sleep(time.Second)
			fmt.Print(".")
		}
	}
}
