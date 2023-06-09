package info

import (
	"fmt"
	"golang.org/x/net/http2"
	"io"
	"net/http"
)

func GetHttpResponse(url string) (string, error) {
	// 创建 HTTP 客户端
	client := &http.Client{
		Transport: &http2.Transport{},
	}

	// 发送 GET 请求
	response, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("请求失败：%s", err.Error())
	}
	defer response.Body.Close()

	// 读取响应内容
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败：%s", err.Error())
	}
	return string(responseData), nil
}
