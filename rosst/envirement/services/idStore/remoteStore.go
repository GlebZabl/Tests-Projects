package idStore

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"Tests-Projects/rosst/envirement/errors"
)

const urlPath = "https://www.uuidgenerator.net/api/version4"

type remoteStore struct {
	client http.Client
}

func (r *remoteStore) GetNewId() (string, error) {
	request, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		return "", Errors.New(err)
	}

	response, err := r.client.Do(request)
	if err != nil {
		return "", Errors.New(err)
	}

	rowId, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", Errors.New(err)
	}

	return strings.Trim(string(rowId), "\r\n"), nil
}

func NewRemoteStore(timeOut int) IdStore {
	defaultTransport := http.DefaultTransport.(*http.Transport)
	defaultTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := http.Client{
		Transport: defaultTransport,
		Timeout:   time.Duration(timeOut) * time.Second,
	}

	return &remoteStore{
		client: client,
	}
}
