package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hfantin/clientgo/db"
	r "github.com/hfantin/clientgo/routes"
)

func main() {
	printLogo()
	// TODO get params in env
	db := db.Connect(getDbEnv())
	router := r.Initialize(db)
	port := getPort()
	log.Println("Serving on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
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
