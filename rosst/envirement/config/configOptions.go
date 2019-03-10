package config

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"Tests-Projects/rosst/envirement/errors"
	"Tests-Projects/rosst/envirement/services/logger"
)

type ConfOptions struct {
	LoggerKey          string    `json:"logger"`
	DbOptions          dbOptions `json:"db_options"`
	Port               string    `json:"port"`
	TimeOutOfRemoteReq int       `json:"time_out_of_remote_req"`
	IsTestEnv          bool      `json:"is_test_env"`
}

type dbOptions struct {
	ConnectionString string `json:"connection_string"`
	DriverName       string `json:"driver_name"`
}

func (c *ConfOptions) Validate() string {
	if !(c.LoggerKey == logger.ConsolePrinterName || c.LoggerKey == logger.FilePrinterName || c.LoggerKey == logger.FullPrinterName) {
		return "invalid config parameter : logger"
	}

	if len(c.Port) < 2 {
		return "invalid config parameter : port"
	}
	_, err := strconv.Atoi(c.Port[1:])
	if string(c.Port[0]) != ":" || err != nil {
		return "invalid config parameter : port"
	}
	return ""
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

	if wrong := config.Validate(); wrong != "" {
		return nil, Errors.NewWithMessage(wrong)
	}
	return config, nil
}
