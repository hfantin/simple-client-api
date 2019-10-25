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
	e "github.com/hfantin/clientgo/eureka"
	r "github.com/hfantin/clientgo/routes"
)

func main() {
	eureka := e.Client{}
	handleSigterm(&eureka)
	go startWebServer()
	eureka.Register()
	go eureka.StartHeartbeat()

	wg := sync.WaitGroup{} // Use a WaitGroup to block main() exit
	wg.Add(1)
	wg.Wait()

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

func handleSigterm(e *e.Client) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		e.Deregister()
		os.Exit(1)
	}()
}

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
