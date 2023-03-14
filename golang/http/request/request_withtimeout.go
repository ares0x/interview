package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	cli := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := cli.Get("http://baidu.com")
	if err != nil {
		fmt.Println("req failed:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
