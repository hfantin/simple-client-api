package utils

import "github.com/twinj/uuid"

func GetUUID() string {
	return uuid.NewV4().String()
}
