package main

import (
	"log"
	"time"

	"github.com/tiger586/go-httpclient"
)

func main() {
	// JSON POST 範例
	payload := map[string]any{
		"name":  "Tiger",
		"email": "tiger@example.com",
	}
	headers := map[string]string{
		"Authorization": "Bearer abc123",
		"Accept":        "application/json",
	}

	body, status, err := httpclient.NewRequest("POST", "https://httpbin.org/post").
		JSON(payload).
		HeadersAdd(headers).
		Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("POST JSON Status:", status)
	httpclient.PrettyPrint(body)

	// Form POST 範例
	form := map[string]string{
		"name":  "Tiger",
		"email": "tiger@example.com",
	}
	body, status, _ = httpclient.NewRequest("POST", "https://httpbin.org/post").
		Form(form).
		Do()
	log.Println("POST Form Status:", status)
	httpclient.PrettyPrint(body)

	// Multipart POST 範例
	fields := map[string]string{
		"name": "Tiger",
	}
	files := map[string]string{
		"file": "./example.txt", // 需本地存在 example.txt
	}
	body, status, _ = httpclient.NewRequest("POST", "https://httpbin.org/post").
		Multipart(fields, files).
		Do()
	log.Println("POST Multipart Status:", status)
	httpclient.PrettyPrint(body)

	log.Println("Set timeout:")
	log.Println("Start:", time.Now().Format("2006-01-02 15:04:05.000"))
	// 自訂 timeout
	// client := httpclient.NewClient(5 * time.Second)
	client := httpclient.NewClient(100 * time.Millisecond)
	resp, status, err := client.NewRequest("POST", "https://httpbin.org/post").
		// resp, status, err := client.NewRequest("POST", "http://localhost:8043/whois").
		JSON(payload).
		HeadersAdd(headers).
		Do()
	if err != nil {
		log.Println("End:", time.Now().Format("2006-01-02 15:04:05.000"))
		log.Fatal(err)
	}
	log.Println("End:", time.Now().Format("2006-01-02 15:04:05.000"))

	log.Println("Status:", status)
	log.Println("Body:", string(resp))
}
