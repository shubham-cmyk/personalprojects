package main

import (
	"fmt"
	"net/http"
)

func (lb *LoadBalancer) ServerProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.GetnextAvailableServer()
	fmt.Printf("Forwarding Request to address %q \n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func (lb *LoadBalancer) GetnextAvailableServer() Server {
	server := lb.servers[lb.rountRobinMethod%len(lb.servers)]

	for !server.IsAlive() {
		lb.rountRobinMethod++
		server = lb.servers[lb.rountRobinMethod%len(lb.servers)]
	}
	lb.rountRobinMethod++
	return server
}

func (s *Server) Address() string {
	return s.address
}

func (s *Server) IsAlive() bool {
	return true
}

func (s *Server) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)

}
