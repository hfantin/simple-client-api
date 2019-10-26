package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("guest", "guest", "test")
	ensureTableExists()
	code := m.Run()
	listClients()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func listClients() {
	a.DB.Exec("select 1 from test.CLI limit 1;")
}

func TestEmptyTable(t *testing.T) {
	req, _ := http.NewRequest("GET", "/clients", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentClient(t *testing.T) {
	req, _ := http.NewRequest("GET", "/client/-1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Client not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Client not found'. Got '%s'", m["error"])
	}
}

func TestCreateClient(t *testing.T) {
	payload := []byte(`{"name":"test client","birthDate":"2019-01-01"}`)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["name"] != "test client" {
		t.Errorf("Expected client name to be 'test client'. Got '%v'", m["name"])
	}
	if m["birthDate"] != "2019-01-01" {
		t.Errorf("Expected user birthDate to be '2019-01-01'. Got '%v'", m["age"])
	}
	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected user ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetUser(t *testing.T) {
	addClient(1)
	req, _ := http.NewRequest("GET", "/client/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func addClient(count int) {
	if count < 1 {
		count = 1
	}
	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO CLI(name, birth_date, email) VALUES('%s', '%s', '%s')", ("Cli " + strconv.Itoa(i+1)), "2019-01-01", "cli@gmail.com")
		a.DB.Exec(statement)
	}
}

func TestUpdateUser(t *testing.T) {
	addClient(1)
	req, _ := http.NewRequest("GET", "/clients/1", nil)
	response := executeRequest(req)
	var originalCli map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalCli)
	payload := []byte(`{"name":"test cli - updated name","birthDate":"2019-01-01"}`)
	req, _ = http.NewRequest("PUT", "/clients/1", bytes.NewBuffer(payload))
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["id"] != originalCli["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalCli["id"], m["id"])
	}
	if m["name"] == originalCli["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalCli["name"], m["name"], m["name"])
	}

	if m["birthDate"] == originalCli["birthDate"] {
		t.Errorf("Expected the birthDate to change from '%v' to '%v'. Got '%v'", originalCli["birthDate"], m["birthDate"], m["birthDate"])
	}
}

func TestDeleteUser(t *testing.T) {
	addClient(1)
	req, _ := http.NewRequest("GET", "/clients/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	req, _ = http.NewRequest("DELETE", "/clients/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	req, _ = http.NewRequest("GET", "/clients/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// func clearTable() {
// 	a.DB.Exec("DELETE FROM CLI")
// 	a.DB.Exec("ALTER TABLE CLI AUTO_INCREMENT = 1")
// }

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS CLI (
	ID int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
	NAME varchar(50) NOT NULL,
	BIRTH_DATE date NOT NULL,
	EMAIL varchar(100) DEFAULT NULL
)`
