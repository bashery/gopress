package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url = "http://localhost:8080"
)

func main() {
	fmt.Println(userAuth("serial"))
	fmt.Println(expiration("serial"))
}

func userAuth(serial string) string {
	resp, err := http.Get(url + "/auth?serial=" + serial)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	user, _ := ioutil.ReadAll(resp.Body)
	return string(user)

}

func expiration(serial string) string {
	resp, err := http.Get(url + "/time?serial=" + serial)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	btime, _ := ioutil.ReadAll(resp.Body)
	return string(btime)

}
