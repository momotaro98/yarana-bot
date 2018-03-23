package main

import (
	"bytes"
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

// HTTPPost posts json contents
func HTTPPost(url string, json []byte) error {
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(json),
	)
	if err != nil {
		return err
	}

	// Setting of Content-Type
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
