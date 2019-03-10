package usersStore

import "Tests-Projects/rosst/envirement/models"

type UsersStore interface {
	AddUser(models.User) error
	GetUserByMail(string) (*models.User, error)
}
