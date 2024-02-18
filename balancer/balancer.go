package balancer

import (
	"fmt"
	"net/http"

	"github.com/thiri-lwin/go-load-balancer/server"
)

type LoadBalancer struct {
	Port            string
	servers         []server.Server
	roundRobinCount int
}

func NewLoadBalancer(port string, servers []server.Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		servers:         servers,
		roundRobinCount: 0,
	}
}

func (lb *LoadBalancer) getNextAvailableServer() server.Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++

	return server
}

func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()

	fmt.Printf("forwarding request to address %q\n", targetServer.Address())

	targetServer.Serve(rw, req)
}
