package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hfantin/clientgo/client"
	"github.com/hfantin/clientgo/server"
	"github.com/hfantin/clientgo/utils"
)

var consul client.Consul

func main() {
	printLogo()
	env := utils.Config()
	handleSigterm()
	consul = client.Consul{}
	consul.Register(env.ServerPort)
	go server.StartWebServer(env)
	wg := sync.WaitGroup{} // Use a WaitGroup to block main() exit
	wg.Add(1)
	wg.Wait()

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

func handleSigterm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Stopping server...")
		consul.Deregister()
		os.Exit(1)
	}()
}
