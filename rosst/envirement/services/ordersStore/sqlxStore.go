package ordersStore

import (
	"database/sql"

	"Tests-Projects/rosst/envirement/errors"
	"Tests-Projects/rosst/envirement/models"

	"github.com/jmoiron/sqlx"
	sqlBuilder "gopkg.in/Masterminds/squirrel.v1"
)

const (
	ordersTable  = "orders"
	idField      = "id"
	ownerIdField = "owner_id"
	infoField    = "info"
	dateField    = "date"
)

type sqlxStore struct {
	db *sqlx.DB
}

func (s *sqlxStore) SaveOrder(o models.Order) error {
	if err := s.db.Ping(); err != nil {
		return Errors.New(err)
	}

	query, args, err := sqlBuilder.Insert(ordersTable).
		Columns(idField, ownerIdField, infoField, dateField).
		Values(o.Id, o.OwnerId, o.Info, o.Date).ToSql()
	if err != nil {
		return Errors.New(err)
	}

	_, err = s.db.Exec(query, args...)
	if err != nil {
		return Errors.New(err)
	}

	return nil
}

func (s *sqlxStore) GetUsersOrders(userId string) ([]*models.Order, error) {
	if err := s.db.Ping(); err != nil {
		return nil, Errors.New(err)
	}

	var result []*models.Order
	query, args, err := sqlBuilder.Select(idField, ownerIdField, infoField, dateField).
		From(ordersTable).
		Where(sqlBuilder.Eq{ownerIdField: userId}).
		OrderBy(dateField).ToSql()
	if err != nil {
		return nil, Errors.New(err)
	}

	err = s.db.Select(&result, query, args...)
	if err == sql.ErrNoRows {
		err = nil
	}

	if err != nil {
		return nil, Errors.New(err)
	}

	return result, nil
}

func NewSqlxStore(db *sqlx.DB) OrderStore {
	return &sqlxStore{
		db: db,
	}
}
