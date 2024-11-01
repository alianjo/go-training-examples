package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, _ := url.Parse(s)
	fmt.Println(u.Host)
	fmt.Println(u.Port())
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
}
