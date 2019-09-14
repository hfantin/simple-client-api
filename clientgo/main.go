package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)


type Client struct {
    id        int32   `json:"id,omitempty"`
    name      string   `json:"firstname,omitempty"`
    birthDate  string   `json:"lastname,omitempty"`
    email   string   `json:"lastname,omitempty"`
}


var clients []Client

func GetClients(w http.ResponseWriter, r *http.Request) {
    clients = append(clients, Client{id: 1, name: "Seu madruga", birthDate: "2006-01-01", email: "m@gmail.com"})
    json.NewEncoder(w).Encode(clients)
}
//func GetPerson(w http.ResponseWriter, r *http.Request) {}
//func CreatePerson(w http.ResponseWriter, r *http.Request) {}
//func DeletePerson(w http.ResponseWriter, r *http.Request) {}

// função principal


func main() {
    println("iniciando...")
    router := mux.NewRouter()
    clients = append(clients, Client{id: 1, name: "Seu madruga", birthDate: "2006-01-01", email: "m@gmail.com"})
    clients = append(clients, Client{id: 2, name: "Kiko", birthDate: "2006-02-01", email: "k@gmail.com"})
    clients = append(clients, Client{id: 3, name: "Chaves", birthDate: "2006-03-01", email: "c@gmail.com"})
    fmt.Printf("%V", clients)
    router.HandleFunc("/clients", GetClients).Methods("GET")
//    router.HandleFunc("/clients/{id}", GetPerson).Methods("GET")
//    router.HandleFunc("/clients/{id}", CreatePerson).Methods("POST")
//	router.HandleFunc("/clients/{id}", DeletePerson).Methods("DELETE") 
	log.Fatal(http.ListenAndServe(":3000", router))
}

