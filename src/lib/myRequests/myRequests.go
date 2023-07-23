package myRequests

import (
	"io/ioutil"
	"net/http"
)

func GetStringResponse(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	response := string(body)

	return response, nil
}

func GetByteResponse(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
