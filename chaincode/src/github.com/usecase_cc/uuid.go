package main

import (
	"github.com/google/uuid"
)

func uuidgen()(string) {
	id := uuid.New()
	return id.String()
}