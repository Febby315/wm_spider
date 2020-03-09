package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//HTTPPost post
func HTTPPost(url string, param map[string]interface{}) string {

	client := &http.Client{}

	pb, _ := json.Marshal(param)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(pb)))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	return string(body)
}

//HTTPGet get
func HTTPGet(url string, param map[string]interface{}) string {
	client := &http.Client{}

	pb, _ := json.Marshal(param)
	req, err := http.NewRequest("GET", url, strings.NewReader(string(pb)))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	return string(body)
}
