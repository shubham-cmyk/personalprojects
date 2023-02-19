package main

import (
	"net/http/httputil"
)

type Server struct {
	address string
	proxy   httputil.ReverseProxy
}

type LoadBalancer struct {
	port             string
	rountRobinMethod int
	servers          []Server
}
