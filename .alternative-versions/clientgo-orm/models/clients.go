package models

import (
	"database/sql"
	"fmt"
)

type Client struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	BirthDate string         `json:"bithDate"`
	Email     sql.NullString `json:"email"`
}

func GetClients() []*Client {
	clients := make([]*Client, 0)
	err := GetDB().Table("CLI").Find(&clients).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return clients
}
