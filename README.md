# go-httpclient [![Go version](https://img.shields.io/github/go-mod/go-version/tiger586/go-httpclient)](https://github.com/tiger586/go-httpclient/blob/main/go.mod)

```
/
├── go.mod
├── client.go        # 核心 Client + DoRequest
├── request.go       # 主要 Request struct + 鏈式呼叫方法
├── multipart.go     # 共用 multipart 工具
├── decode.go        # JSON Decode / PrettyPrint helper
```

一個輕量級的 Go HTTP 客戶端，支援 JSON、Form 以及 Multipart 請求。  
適合 restful web service client，提供鏈式呼叫 API，讓 HTTP 請求寫法乾淨、統一。

---

## 功能特色

- 鏈式呼叫 Request
- 支援 JSON / Form / Multipart
- 可用於 GET / POST / PUT / PATCH / DELETE
- 簡單新增自訂 Headers
- 內建 JSON PrettyPrint 輔助函數

---

## 安裝

```bash
go get github.com/tiger586/go-httpclient
```

## 使用範例
1. POST JSON
```go
payload := map[string]any{
    "name": "Tiger",
    "email": "tiger@example.com",
}

headers := map[string]string{
    "Authorization": "Bearer abc123",
    "Accept": "application/json",
}

body, status, err := httpclient.NewRequest("POST", "https://httpbin.org/post").
    JSON(payload).
    HeadersAdd(headers).
    Do()
if err != nil {
    log.Fatal(err)
}

httpclient.PrettyPrint(body)
```

2. POST Form
```go
form := map[string]string{
    "name": "Tiger",
    "email": "tiger@example.com",
}

body, status, _ = httpclient.NewRequest("POST", "https://httpbin.org/post").
    Form(form).
    Do()

httpclient.PrettyPrint(body)
```

3. POST Multipart (fields + files 含欄位與檔案)
```go
fields := map[string]string{
    "name": "Tiger",
}

files := map[string]string{
    "file": "./test.png",
}

body, status, _ = httpclient.NewRequest("POST", "https://httpbin.org/post").
    Multipart(fields, files).
    Do()

httpclient.PrettyPrint(body)
```

## 輔助函數
- PrettyPrint(body []byte)  
將 JSON 回傳格式化輸出，如果不是 JSON，則原樣輸出文字。
- GetJSONField(body []byte, key string) string  
直接從返回的 body 取值，支援巢狀 key，例如 "data.user.name"  
如果 body 不是 JSON 或 key 不存在，回傳空字串。

## 授權

MIT License