package client

import (
	"fmt"
	"log"
	"os"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

func RegisterServiceWithConsul(port string) {
	log.Println("register service with consul", port)
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "clients-" + port
	registration.Name = "clients"
	address := hostname()
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
	consul.Agent().ServiceRegister(registration)
}

// func getIp() string {
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		return ""
// 	}
// 	for _, address := range addrs {
// 		// check the address type and if it is not a loopback the display it
// 		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
// 			if ipnet.IP.To4() != nil {
// 				return ipnet.IP.String()
// 			}
// 		}
// 	}
// 	return ""
// }

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}