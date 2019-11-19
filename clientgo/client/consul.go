package client

import (
	"fmt"
	"log"
	"net"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

type Consul struct {
	consul *consulapi.Client
	port   string
}

func newClient() *consulapi.Client {
	log.Println("creating new consul client")
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	return consul
}

func (c *Consul) Register(port string) {
	log.Println("register service with consul", port)
	c.consul = newClient()
	c.port = port
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "clients-" + port
	registration.Name = "clients"
	address := getIp()
	registration.Address = address
	log.Println("register service with consul - address", address)
	porta, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln(err)
	}

	registration.Port = porta
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/actuator/health",
		address, port)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	c.consul.Agent().ServiceRegister(registration)
}

func (c *Consul) Deregister() {
	log.Println("deregister service ", c.port)
	c.consul.Agent().ServiceDeregister("clients-" + c.port)
}

func getIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// func hostname() string {
// 	hn, err := os.Hostname()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return hn
// }
