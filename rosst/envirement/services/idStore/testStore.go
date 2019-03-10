package idStore

import (
	"github.com/satori/go.uuid"
	"test_project/envirement/errors"
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
