package model

import (
	"database/sql"
	"fmt"
)

type Client struct {
	ID        int             `json:"id"`
	Name      *string         `json:"name"`
	BirthDate *string         `json:"bithDate"`
	Email     *sql.NullString `json:"email"`
}

func (c *Client) GetClient(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name, birth_date, email FROM CLI WHERE id=%d", c.ID)
	return db.QueryRow(statement).Scan(&c.Name, &c.BirthDate, &c.Email)
}

func GetAllClients(db *sql.DB) ([]Client, error) {
	statement := fmt.Sprintf("SELECT id, name, birth_date, email FROM CLI")
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	clients := []Client{}
	for rows.Next() {
		var c Client
		if err := rows.Scan(&c.ID, &c.Name, &c.BirthDate, &c.Email); err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}
	return clients, nil
}

func GetClients(db *sql.DB, start, count int) ([]Client, error) {
	statement := fmt.Sprintf("SELECT id, name, birth_date, email FROM CLI LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	clients := []Client{}
	for rows.Next() {
		var c Client
		if err := rows.Scan(&c.ID, &c.Name, &c.BirthDate, &c.Email); err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}
	return clients, nil
}
