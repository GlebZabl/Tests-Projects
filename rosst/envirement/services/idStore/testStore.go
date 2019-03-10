package idStore

import (
	"Tests-Projects/rosst/envirement/errors"
	"github.com/satori/go.uuid"
)

type testStore struct{}

func (t *testStore) GetNewId() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", Errors.New(err)
	}
	return id.String(), nil
}

func NewTestStore() IdStore {
	return new(testStore)
}
