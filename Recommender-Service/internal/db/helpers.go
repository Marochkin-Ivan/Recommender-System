package db

import "github.com/google/uuid"

func GenerateUID() string {
	return uuid.New().String()
}
