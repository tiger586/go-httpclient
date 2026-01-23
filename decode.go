package httpclient

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 將 []byte 嘗試轉成 JSON Pretty Print
func PrettyPrint(body []byte) {
	var v any
	if err := json.Unmarshal(body, &v); err != nil {
		fmt.Println(string(body))
		return
	}

	out, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(out))
}

// GetJSONField 支援巢狀 key，例如 "data.user.name"
// 如果 body 不是 JSON 或 key 不存在，回傳空字串
func GetJSONField(body []byte, key string) string {
	var data any
	if err := json.Unmarshal(body, &data); err != nil {
		return ""
	}

	parts := strings.Split(key, ".")
	var current any = data

	for _, k := range parts {
		m, ok := current.(map[string]any)
		if !ok {
			return ""
		}

		v, exists := m[k]
		if !exists {
			return ""
		}

		current = v
	}

	// 最後一層斷言成 string
	if s, ok := current.(string); ok {
		return s
	}

	return ""
}
