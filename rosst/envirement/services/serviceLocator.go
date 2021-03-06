package services

import (
	"Tests-Projects/rosst/envirement/config"
	"Tests-Projects/rosst/envirement/errors"
	"Tests-Projects/rosst/envirement/services/idStore"
	"Tests-Projects/rosst/envirement/services/logger"
	"Tests-Projects/rosst/envirement/services/ordersStore"
	"Tests-Projects/rosst/envirement/services/usersStore"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const defaultLoggerKey = logger.ConsolePrinterName

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

func InitTestEnv(config config.ConfOptions) error {
	iStore := idStore.NewTestStore()
	uStore := usersStore.NewTestStore()
	oStore := ordersStore.NewTestStore()
	log := logger.NewLogger(config.LoggerKey)
	locator = ServiceLocator{
		os:  oStore,
		us:  uStore,
		is:  iStore,
		l:   log,
		cfg: config,
	}

	return nil
}
