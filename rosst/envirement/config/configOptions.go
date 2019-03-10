package config

import (
	"encoding/json"
	"io/ioutil"
	"test_project/envirement/errors"
)

type ConfOptions struct {
	LoggerKey          string    `json:"logger"`
	DbOptions          dbOptions `json:"db_options"`
	Port               string    `json:"port"`
	TimeOutOfRemoteReq int       `json:"time_out_of_remote_req"`
}

type dbOptions struct {
	ConnectionString string `json:"connection_string"`
	DriverName       string `json:"driver_name"`
}

func LoadConfig(path string) (*ConfOptions, error) {
	rawData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, Errors.New(err)
	}

	config := new(ConfOptions)
	err = json.Unmarshal(rawData, config)
	if err != nil {
		return nil, Errors.New(err)
	}

	return config, nil
}
