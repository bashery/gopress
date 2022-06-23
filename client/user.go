package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(userAuth())
	fmt.Println(expiration())
	fmt.Println(serial())

}

func userAuth() string {
	resp, err := http.Get("http://localhost:8080/auth?serial=" + "serial")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	user, _ := ioutil.ReadAll(resp.Body)
	return string(user)

}

func expiration() string {
	resp, err := http.Get("http://localhost:8080/expiration?time=" + "serial")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	btime, _ := ioutil.ReadAll(resp.Body)
	return string(btime)

}
