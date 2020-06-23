package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

type Proxy struct {
	Proxy      string `json:"proxy"`
	FailCount  int    `json:"fail_count"`
	Region     string `json:"region"`
	Type       string `json:"type "`
	Source     string `json:"source"`
	CheckCount int    `json:"check_count"`
	LastStatus int    `json:"last_status"`
	LastTime   string `json:"last_time"`
}

// 过滤 可用 代理IP
func main() {
	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// 	var proxies []*url.URL = []*url.URL{
	// 	&url.URL{Host: "180.149.144.176:80"},
	// 	&url.URL{Host: "222.74.202.226:9999"},
	// }
	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("http://139.217.110.76:3128")
	// rp, err := proxy.RoundRobinProxySwitcher("http://180.149.144.176:80")

	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	// Fetch httpbin.org/ip five times
	// c.Visit("https://httpbin.org/ip")
	err1 := c.Visit("https://httpbin.org/ip")
	if err1 != nil {
		fmt.Println(err1.Error())
	} else {
		fmt.Println("111")
	}

}
