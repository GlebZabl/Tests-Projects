package usersStore

import (
	"database/sql"

	"test_project/envirement/errors"
	"test_project/envirement/models"

	"github.com/jmoiron/sqlx"
	sqlBuilder "gopkg.in/Masterminds/squirrel.v1"
)

const (
	usersTable = "users"
	idField    = "id"
	mailField  = "mail"
)

type sqlxStore struct {
	db *sqlx.DB
}

func (s *sqlxStore) AddUser(u models.User) error {
	if err := s.db.Ping(); err != nil {
		return Errors.New(err)
	}

	query, args, err := sqlBuilder.Insert(usersTable).
		Columns(idField, mailField).
		Values(u.Id, u.Mail).ToSql()
	if err != nil {
		return Errors.New(err)
	}

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return Errors.New(err)
	}

	return nil
}

func (s *sqlxStore) GetUserByMail(mail string) (*models.User, error) {
	if err := s.db.Ping(); err != nil {
		return nil, Errors.New(err)
	}

	result := new(models.User)
	query, args, err := sqlBuilder.Select(idField, mailField).
		From(usersTable).
		Where(sqlBuilder.Eq{mailField: mail}).ToSql()
	if err != nil {
		return nil, Errors.New(err)
	}

	err = s.db.QueryRowx(query, args...).StructScan(result)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, Errors.New(err)
	}

	return result, nil
}

func NewSqlxStore(db *sqlx.DB) UsersStore {
	return &sqlxStore{
		db: db,
	}
}
