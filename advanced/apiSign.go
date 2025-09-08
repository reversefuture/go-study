package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// 生成签名
func generateSignature(params map[string]string, secretKey string) string {
	// 1. 提取参数名并排序
	var keys []string
	for k := range params {
		if k != "sign" { // 排除 sign 参数本身
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// 2. 拼接待签名字符串
	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, params[k]))
	}
	concatenated := strings.Join(parts, "&")

	// 3. 使用 HMAC-SHA256 签名
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(concatenated))
	signature := hex.EncodeToString(h.Sum(nil))

	return signature
}

func mainAPI() {
	params := map[string]string{
		"timestamp": "1712345678",
		"nonce":     "abc123",
		"action":    "getUser",
		"userId":    "1001",
	}

	secretKey := "your-secret-key-here"

	sign := generateSignature(params, secretKey)
	params["sign"] = sign

	fmt.Printf("请求参数: %+v\n", params)
	fmt.Printf("生成的签名: %s\n", sign)
}
