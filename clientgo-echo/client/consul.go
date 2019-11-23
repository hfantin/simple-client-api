package client

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

type Consul struct {
	consul *consulapi.Client
	port   string
}

func newClient() *consulapi.Client {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	return consul
}

func (c *Consul) Register(port string) {
	c.consul = newClient()
	c.port = port
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "clients-" + port
	registration.Name = "clients"
	address := getIp() //hostname()
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

func (c *Consul) service(service, tag string) ([]*consulapi.ServiceEntry, *consulapi.QueryMeta, error) {
	passingOnly := true
	addrs, meta, err := c.consul.Health().Service(service, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, nil, fmt.Errorf("service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, nil, err
	}
	return addrs, meta, nil
}

func (c *Consul) Resolve(service, tag string) (string, error) {
	entries, _, err := c.service(service, tag)
	if err != nil {
		return "", err
	}
	index := 0
	if len(entries) > 1 {
		rand.Seed(time.Now().Unix())
		index = rand.Intn(len(entries))
	}
	entry := entries[index]
	uri := fmt.Sprintf("http://%s:%d/", entry.Service.Address, entry.Service.Port)
	return uri, nil
}
