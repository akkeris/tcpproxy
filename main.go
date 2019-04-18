package main

import (
	"github.com/google/tcpproxy"
	"log"
	"os"
)

func main() {
	var p tcpproxy.Proxy
	p.AddRoute(":"+os.Getenv("PORT"), tcpproxy.To(os.Getenv("PROXY_ENDPOINT")))
	log.Fatal(p.Run())
}
