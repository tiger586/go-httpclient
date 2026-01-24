package httpclient

import (
	"net/http"
	"time"
)

type Client struct {
	http *http.Client
}

// 預設 timeout
var defaultClient = &Client{
	http: &http.Client{Timeout: 10 * time.Second}, // 預設 timeout
}

// 建構函數，自訂 timeout
func NewClient(timeout time.Duration) *Client {
	return &Client{
		http: &http.Client{Timeout: timeout},
	}
}
