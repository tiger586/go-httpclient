package main

import (
	"fmt"
	"time"

	"github.com/tiger586/go-httpclient"
)

func main() {
	// 建立 client，自訂 timeout = 2 秒
	client := httpclient.NewClient(2 * time.Second)

	fmt.Println("start request at:", time.Now().Format(time.RFC3339))

	// 非同步執行 request
	go func() {
		body, status, err := client.
			NewRequest("POST", "https://httpbin.org/post").
			// body, status, err := client.
			// NewRequest("POST", "http://localhost:8043/whois").
			JSON(map[string]string{
				"domain": "example.com",
			}).
			Do()

		fmt.Println("---- async result ----")
		fmt.Println("time:", time.Now().Format(time.RFC3339))
		fmt.Println("status:", status)
		fmt.Println("err:", err)
		fmt.Println("body:", string(body))
	}()

	// 模擬主程式繼續做事
	for i := 1; i <= 5; i++ {
		fmt.Println("main doing other work...", i)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("main finished")
}
