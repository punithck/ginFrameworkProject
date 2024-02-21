package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://localhost:8081/videos"
	method := "GET"

	// First create a client from http library
	client := &http.Client{}
	// Create a new request
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	// Add header to the request, adding passwords
	req.Header.Add("Authorization", "Basic cHU6bmk=")

	// Calling the function
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Closing the res at the dn
	defer res.Body.Close()

	// Reading all data
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
