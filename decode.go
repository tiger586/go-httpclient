package httpclient

import (
	"encoding/json"
	"fmt"
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
