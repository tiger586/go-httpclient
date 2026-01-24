package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"maps"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	client  *Client
	Method  string
	URL     string
	Headers map[string]string
	Body    io.Reader
}

// 建立 Request
//
// method = GET / POST / PUT / PATCH / DELETE
// NewRequest 保留舊用法，使用預設 client（10 秒 timeout）
//
//	func NewRequest(method, url string) *Request {
//		method = strings.ToUpper(method) // ✅ 自動轉大寫
//		return &Request{
//			Method:  method,
//			URL:     url,
//			Headers: map[string]string{},
//		}
//	}
func NewRequest(method, url string) *Request {
	return defaultClient.NewRequest(method, url)
}

func (c *Client) NewRequest(method, url string) *Request {
	method = strings.ToUpper(method)
	return &Request{
		client:  c, // 指向你自訂的 client
		Method:  method,
		URL:     url,
		Headers: map[string]string{},
	}
}

// 設定 JSON body
func (r *Request) JSON(payload any) *Request {
	data, _ := json.Marshal(payload)
	r.Body = bytes.NewBuffer(data)
	r.Headers["Content-Type"] = "application/json"
	return r
}

// 設定 form-urlencoded body
func (r *Request) Form(form map[string]string) *Request {
	values := url.Values{}
	for k, v := range form {
		values.Set(k, v)
	}
	r.Body = strings.NewReader(values.Encode())
	r.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	return r
}

// 設定 multipart body
func (r *Request) Multipart(fields map[string]string, files map[string]string) *Request {
	body, contentType, _ := NewMultipart(fields, files)
	r.Body = body
	r.Headers["Content-Type"] = contentType
	return r
}

// 設定自訂 headers
func (r *Request) HeadersAdd(headers map[string]string) *Request {
	// for k, v := range headers {
	// 	r.Headers[k] = v
	// }
	maps.Copy(r.Headers, headers)
	return r
}

// 執行 request，回傳 body + status + error
func (r *Request) Do() ([]byte, int, error) {
	req, err := http.NewRequest(r.Method, r.URL, r.Body)
	if err != nil {
		return nil, 0, err
	}
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	// resp, err := defaultClient.Do(req)
	resp, err := r.client.http.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return body, resp.StatusCode, nil
}
