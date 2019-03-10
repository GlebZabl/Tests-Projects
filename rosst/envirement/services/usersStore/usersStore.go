package usersStore

import "test_project/envirement/models"

type UsersStore interface {
	AddUser(models.User) error
	GetUserByMail(string) (*models.User, error)
}
