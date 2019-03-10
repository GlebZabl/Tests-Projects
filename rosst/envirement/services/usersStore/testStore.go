package usersStore

import (
	"sync"

	"Tests-Projects/rosst/envirement/errors"
	"Tests-Projects/rosst/envirement/models"
)

var users safeMap

func init() {
	users.data = make(map[string]models.User)
}

type safeMap struct {
	sync.Mutex
	data map[string]models.User
}

func NewTestStore() UsersStore {
	return new(TestStore)
}

type TestStore struct{}

func (t *TestStore) AddUser(u models.User) error {
	users.Lock()
	defer users.Unlock()

	if _, ok := users.data[u.Id]; ok {
		return Errors.NewWithMessage("users with this Id is already exist")
	}

	users.data[u.Id] = u
	return nil
}

func (t *TestStore) GetUserByMail(mail string) (*models.User, error) {
	users.Lock()
	defer users.Unlock()

	for _, u := range users.data {
		if u.Mail == mail {
			return &u, nil
		}
	}

	return nil, nil
}

//for tests only
func (t *TestStore) GetData() map[string]models.User {
	users.Lock()
	defer users.Unlock()

	return users.data
}

func (t *TestStore) SetData(data map[string]models.User) {
	users.Lock()
	defer users.Unlock()

	users.data = data
}
