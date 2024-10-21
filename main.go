package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
)

// appList 函数用于获取appforge智能体列表，通过日志打印出appforge智能体列表
func appList(url string) (a []byte, err error) {
	// 创建一个自定义的 Transport 层，跳过 TLS 证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-cc-authdata", "eyJsb2dpbklkIjoiMTg0MTc4MTY2Mjg0OTk3Nzc1NSIsImxvZ2luQ29kZSI6Imd1b2R4MSIsImNsb3VkTG9naW5JZCI6IjE4NDE3ODE2NjI4NDk5Nzc3NTMiLCJjbG91ZExvZ2luQ29kZSI6Imd1b2R4MSIsInBpblR5cGUiOjAsInRlbmFudElkIjoiMTg0MTc4MTY2Mjg0OTk3Nzc1MyIsImRlcGFydG1lbnRJZCI6IjAiLCJ0ZW5hbnRVc2VyVHlwZSI6Mn0=")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	//defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	return body, nil
}

func main() {
	url := "http://10.252.20.1:31642/console/api/apps"
	resp, err := appList(url)
	if err != nil {
		log.Fatalf("获取appforge智能体列表失败: %v", err)
	}

	// 打印appforge智能体列表
	log.Println("appforge智能体列表:", string(resp))
}
