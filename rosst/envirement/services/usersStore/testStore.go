package usersStore

import (
	"test_project/envirement/errors"
	"test_project/envirement/models"
)

var users = map[string]models.User{}

func init() {
	users = make(map[string]models.User)
}

type TestStore struct{}

func (t *TestStore) AddUser(u models.User) error {
	if _, ok := users[u.Id]; ok {
		return Errors.NewWithMessage("users with this Id is already exist")
	}

	users[u.Id] = u
	return nil
}

func (t *TestStore) GetUserByMail(mail string) (*models.User, error) {
	for _, u := range users {
		if u.Mail == mail {
			return &u, nil
		}
	}

	return nil, nil
}

func (t *TestStore) GetData() map[string]models.User {
	return users
}

func (t *TestStore) SetData(data map[string]models.User) {
	users = data
}

func NewTestStore() UsersStore {
	return new(TestStore)
}
