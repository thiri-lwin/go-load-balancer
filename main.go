package main

import (
	"fmt"
	"net/http"

	"github.com/thiri-lwin/go-load-balancer/balancer"
	"github.com/thiri-lwin/go-load-balancer/server"
)

func main() {
	servers := []server.Server{
		server.NewSimpleServer("http://localhost:8080"),
		server.NewSimpleServer("http://localhost:8081"),
		server.NewSimpleServer("http://localhost:8082"),
	}

	loadBalancer := balancer.NewLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		loadBalancer.ServeProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", loadBalancer.Port)
	http.ListenAndServe(":"+loadBalancer.Port, nil)
}
