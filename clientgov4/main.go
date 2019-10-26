package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hfantin/clientgo/db"
	r "github.com/hfantin/clientgo/routes"
	"github.com/joho/godotenv"
)

func main() {
	handleSigterm()
	go startWebServer()
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

// func startConsul(addrs []string, ttl time.Duration) (*Service, error) {
// 	c, err := consul.NewClient(consul.DefaultConfig())
// 	if err != nil {
// 		return nil, err
// 	}
// 	s.ConsulAgent = c.Agent()

// 	serviceDef := &consul.AgentServiceRegistration{
// 		Name: s.Name,
// 		Check: &consul.AgentServiceCheck{
// 			TTL: s.TTL.String(),
// 		},
// 	}

// 	if err := s.ConsulAgent.ServiceRegister(serviceDef); err != nil {
// 		return nil, err
// 	}

// }

// func (s *Service) UpdateTTL(check func() (bool, error)) {
// 	ticker := time.NewTicker(s.TTL / 2)
// 	for range ticker.C {
// 		s.update(check)
// 	}
// }

// TODO put in banner.txt file
func printLogo() {
	fmt.Println("    ______               _      _____  _____ ")
	fmt.Println("   / ____/ (_)__  ____  / /_   / ____/ __  /")
	fmt.Println("  / /   / / / _  / __  / __/  / / __/ / / /")
	fmt.Println(" / /___/ / /  __/ / / / /_   / /_/ / /_/ / ")
	fmt.Println("/_____/_/_/____/_/ /_/ __/   _____/ ____/ ")
	fmt.Println("")
}

func handleSigterm() {
	log.Println("Stopping server...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
}

func startWebServer() {
	printLogo()
	// Default values
	port := "3000"
	host := "localhost"
	user := "guest"
	pass := "guest"
	name := "test"
	// Load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using default values....")
	}
	port = os.Getenv("SERVER_PORT")
	host = os.Getenv("DB_HOST")
	user = os.Getenv("DB_USER")
	pass = os.Getenv("DB_PASS")
	name = os.Getenv("DB_NAME")
	log.Println("db_host=", host)
	// initialize db connection
	db := db.Connect(host, user, pass, name)
	// initialize router
	router := r.Initialize(db)
	log.Println("Serving on port", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port", port)
		log.Println("Error: " + err.Error())
	}
}
