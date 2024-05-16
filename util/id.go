package util

import "github.com/google/uuid"

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func NewUUIDPointer() *uuid.UUID {
	uuid := GenerateUUID()
	return &uuid
}
