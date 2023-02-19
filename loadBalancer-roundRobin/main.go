package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	servers := []Server{
		*NewServer("https://www.facebook.com"),
		*NewServer("https://www.google.com"),
		*NewServer("https://www.bing.com"),
	}

	lb := NewLoadBalancer("8000", servers)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServerProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Serving Request at 'localhost;%s' \n", lb.port)
	err := http.ListenAndServe(":"+lb.port, nil)

	if err != nil {
		fmt.Printf("Error occured while Serving %v", err.Error())
		os.Exit(1)
	}

}
