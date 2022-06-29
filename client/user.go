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
	userAuth("serial")
	expiration("serial")
}

// check if bot is signup by serial
func userAuth(serial string) {
	resp, err := http.Get(url + "/auth?serial=" + serial)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	user, _ := ioutil.ReadAll(resp.Body)

	println(string(user))

}

func expiration(serial string) {
	resp, err := http.Get(url + "/time?serial=" + serial)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	btime, _ := ioutil.ReadAll(resp.Body)
	println(string(btime))

}
