package server

import (
	"io/ioutil"
	"net/http"
)

// Gets data from URL and returns the body of data as []byte and an error
func Harvest(url string) ([]byte, error) {
	res, err := http.Get(url) //get contents of API
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body) //read body of contents
	if err != nil {
		return []byte{}, err
	}

	return response, err
}
