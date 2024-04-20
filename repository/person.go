package repository

import (
	"errors"
	"github.com/lucasDaru/testify_in_practice/repositorytest"

	"github.com/google/uuid"
	"github.com/lucasDaru/testify_in_practice/model"
)

// FindPersonByID simulates searching for a person in the database by ID
func FindPersonByID(personID string) (*model.Person, error) {
	// Simulating a database
	db := repositorytest.TestPersons

	if !IsUUID(personID) {
		return nil, errors.New("invalid input syntax for type uuid: " + personID)
	}

	person, ok := db[personID]
	if !ok {
		return nil, errors.New("not found")
	}

	return &person, nil
}

// IsUUID checks if the given string is a valid UUID
func IsUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
