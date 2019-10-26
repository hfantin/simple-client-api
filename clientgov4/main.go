package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	consul "github.com/hashicorp/consul/api"
	"github.com/hfantin/clientgo/db"
	r "github.com/hfantin/clientgo/routes"
)

func main() {
	// eureka := e.Client{}
	// handleSigterm(&eureka)
	go startWebServer()
	// eureka.Register()
	// go eureka.StartHeartbeat()

	wg := sync.WaitGroup{} // Use a WaitGroup to block main() exit
	wg.Add(1)
	wg.Wait()

}

// func startConsul() {
// 	// Service discovery domain. In this example we use Consul.
// 	var client consulsd.Client
// 	{
// 		consulConfig := api.DefaultConfig()
// 		if len(*consulAddr) > 0 {
// 			consulConfig.Address = *consulAddr
// 		}
// 		consulClient, err := api.NewClient(consulConfig)
// 		if err != nil {
// 			logger.Log("err", err)
// 			os.Exit(1)
// 		}
// 		client = consulsd.NewClient(consulClient)
// 	}
// }

func startConsul(addrs []string, ttl time.Duration) (*Service, error) {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, err
	}
	s.ConsulAgent = c.Agent()

	serviceDef := &consul.AgentServiceRegistration{
		Name: s.Name,
		Check: &consul.AgentServiceCheck{
			TTL: s.TTL.String(),
		},
	}

	if err := s.ConsulAgent.ServiceRegister(serviceDef); err != nil {
		return nil, err
	}

}

func (s *Service) UpdateTTL(check func() (bool, error)) {
	ticker := time.NewTicker(s.TTL / 2)
	for range ticker.C {
		s.update(check)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}

func getDbEnv() (string, string, string) {
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "guest"
	}
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = "guest"
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "test"
	}
	return user, pass, name
}

// TODO put in banner.txt file
func printLogo() {
	fmt.Println("    ______               _      _____  _____ ")
	fmt.Println("   / ____/ (_)__  ____  / /_   / ____/ __  /")
	fmt.Println("  / /   / / / _  / __  / __/  / / __/ / / /")
	fmt.Println(" / /___/ / /  __/ / / / /_   / /_/ / /_/ / ")
	fmt.Println("/_____/_/_/____/_/ /_/ __/   _____/ ____/ ")
	fmt.Println("")
}

// func handleSigterm(e *e.Client) {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)
// 	signal.Notify(c, syscall.SIGTERM)
// 	go func() {
// 		<-c
// 		e.Deregister()
// 		os.Exit(1)
// 	}()
// }

func startWebServer() {

	printLogo()
	// TODO get params in env
	db := db.Connect(getDbEnv())
	router := r.Initialize(db)
	port := getPort()
	log.Println("Serving on port", port)
	// log.Fatal(http.ListenAndServe(":"+port, router))
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port 8080")
		log.Println("Error: " + err.Error())
	}
}
