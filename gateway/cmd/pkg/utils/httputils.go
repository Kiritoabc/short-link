package utils

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func ForwardRequest(targetURL *url.URL, originalReq *http.Request) []byte {
	// 解析目标 URL
	client := &http.Client{}

	// 构建新的请求
	newReq, err := http.NewRequest(originalReq.Method, targetURL.String()+originalReq.RequestURI, originalReq.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 复制请求头
	newReq.Header = originalReq.Header

	// 发送请求
	resp, err := client.Do(newReq)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 处理响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
