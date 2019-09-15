package main

import (
	"database/sql"
	"fmt"
)

type client struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	BirthDate string         `json:"bithDate"`
	Email     sql.NullString `json:"email"`
}

func (u *client) getClient(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name, birth_date, email FROM CLI WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Name, &u.BirthDate, &u.Email)
}

func (u *client) updateClient(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE CLI SET name='%s', birth_date='%s', email='%s' WHERE id=%d", u.Name, u.BirthDate, u.Email, u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *client) deleteClient(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM CLI WHERE id=%d", u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *client) createClient(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO CLI(name, birth_date, email) VALUES('%s', '%s', '%s')", u.Name, u.BirthDate, u.Email)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

func getAllClients(db *sql.DB) ([]client, error) {
	statement := fmt.Sprintf("SELECT id, name, birth_date, email FROM CLI")
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	clients := []client{}
	for rows.Next() {
		var u client
		if err := rows.Scan(&u.ID, &u.Name, &u.BirthDate, &u.Email); err != nil {
			return nil, err
		}
		clients = append(clients, u)
	}
	return clients, nil
}

func getClients(db *sql.DB, start, count int) ([]client, error) {
	statement := fmt.Sprintf("SELECT id, name, birth_date, email FROM CLI LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	clients := []client{}
	for rows.Next() {
		var u client
		if err := rows.Scan(&u.ID, &u.Name, &u.BirthDate, &u.Email); err != nil {
			return nil, err
		}
		clients = append(clients, u)
	}
	return clients, nil
}
