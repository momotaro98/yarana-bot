package main

import (
	"io/ioutil"
	"net/http"
)

// AssembleURLWithParam makes url with parameter
func AssembleURLWithParam(baseURL string, key string, val string) string {
	return baseURL + "?" + key + "=" + val
}

// HTTPGet curls with arg url
func HTTPGet(url string) (buf []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return buf, err
}
