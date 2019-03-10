package services

import (
	"test_project/envirement/config"
	"test_project/envirement/errors"
	"test_project/envirement/services/idStore"
	"test_project/envirement/services/logger"
	"test_project/envirement/services/ordersStore"
	"test_project/envirement/services/usersStore"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const defaultLoggerKey = "console"

var (
	locator = ServiceLocator{}
)

type ServiceLocator struct {
	cfg config.ConfOptions
	l   logger.Logger
	os  ordersStore.OrderStore
	us  usersStore.UsersStore
	is  idStore.IdStore
}

func (s *ServiceLocator) Logger() logger.Logger {
	if locator.l == nil {
		return logger.NewLogger(defaultLoggerKey)
	}
	return s.l
}

func (s *ServiceLocator) OrdersStore() ordersStore.OrderStore {
	return s.os
}

func (s *ServiceLocator) UsersStore() usersStore.UsersStore {
	return s.us
}

func (s *ServiceLocator) IdStore() idStore.IdStore {
	return s.is
}

//farther is functions that can be called only from endpoints
func GetEnvironment() *ServiceLocator {
	return &locator
}

func Logger() logger.Logger {
	if locator.l == nil {
		return logger.NewLogger(defaultLoggerKey)
	}
	return locator.l
}

func Config() config.ConfOptions {
	return locator.cfg
}

func InitServices(config config.ConfOptions) error {
	l := logger.NewLogger(config.LoggerKey)
	is := idStore.NewRemoteStore(config.TimeOutOfRemoteReq)
	db, err := sqlx.Connect(config.DbOptions.DriverName, config.DbOptions.ConnectionString)
	if err != nil {
		return Errors.New(err)
	}

	os := ordersStore.NewSqlxStore(db)
	us := usersStore.NewSqlxStore(db)
	locator = ServiceLocator{
		l:   l,
		is:  is,
		os:  os,
		us:  us,
		cfg: config,
	}
	return nil
}

//for tests only!
func SetTestEnv() {
	iStore := idStore.NewTestStore()
	uStore := usersStore.NewTestStore()
	oStore := ordersStore.NewTestStore()
	locator = ServiceLocator{
		os: oStore,
		us: uStore,
		is: iStore,
	}
}
