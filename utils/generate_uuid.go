package utils

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

func GenerateUUID() (string, error) {
	randId, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return "", err
	}

	id := randId.String()
	if id == "" {
		return "", errors.New("Error generating string UUID")
	}

	return id, nil
}
