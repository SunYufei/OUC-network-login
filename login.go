package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// flag
	var username, password string
	flag.StringVar(&username, "u", "", "Username")
	flag.StringVar(&password, "p", "", "Password")
	flag.Parse()
	// request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	url := fmt.Sprintf("https://10.81.2.6/drcom/login?callback=dr1003&DDDDD=%s&upass=%s&0MKKey&R1=0&R2=&R3=0&R6=0&para=00&v6ip=&terminal_type=1&lang=zh-cn&jsVersion=4.1",
		username, password)
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("登陆失败")
		fmt.Println(err)
	} else {
		fmt.Println("登录成功")
	}
	fmt.Println("5秒后自动退出...")
	time.Sleep(5 * time.Second)
}
