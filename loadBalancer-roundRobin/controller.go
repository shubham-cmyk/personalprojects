package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"
)

func NewServer(addr string) *Server {

	url, err := url.Parse(addr)
	if err != nil {
		fmt.Printf("error occured details are :%v \n", err)
		os.Exit(1)
	}

	return &Server{
		address: addr,
		proxy:   *httputil.NewSingleHostReverseProxy(url),
	}
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:             port,
		rountRobinMethod: 0,
		servers:          servers,
	}
}
