package repository_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lucasDaru/testify_in_practice/repository"
	"github.com/stretchr/testify/assert"
)

func TestFindPersonByID(t *testing.T) {
	cases := map[string]struct {
		personID      string
		errorExpected error
	}{
		"success": {
			personID:      "c834d682-78d3-11ed-9315-66c94c3ada1b",
			errorExpected: nil,
		},
		"invalid-person-id": {
			personID:      "999",
			errorExpected: fmt.Errorf("invalid input syntax for type uuid: 999"),
		},
		"person-id-not-found": {
			personID:      "8bb534e0-90dd-4d52-b303-650ee3f343ae",
			errorExpected: errors.New("not found"),
		},
	}

	for testName, testCase := range cases {
		t.Run(testName, func(t *testing.T) {
			person, err := repository.FindPersonByID(testCase.personID)

			if testCase.errorExpected != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, testCase.errorExpected.Error())
				assert.Nil(t, person)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, person)
				assert.Equal(t, testCase.personID, person.ID)
			}
		})
	}
}
