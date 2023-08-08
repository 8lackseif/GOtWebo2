package myRequests

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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

func PostQuery(link string, query string, variables string) (string, error) {

	data := url.Values{
		"variables": {variables},
		"query":     {query},
	}
	resp, err := http.PostForm(link, data)
	if err != nil {
		return "error post request", err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return "error converting http response to string", err
	}

	if resp.StatusCode == 404 {
		return "make sure the options are correct", nil
	}

	return string(b), nil
}
