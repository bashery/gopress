package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
	resp, err := http.Get(url + "/info?serial=" + serial)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	user, _ := ioutil.ReadAll(resp.Body)
	if len(user) < 5 {
		fmt.Println("no auth ")
		os.Exit(0)
	}
	println(string(user))
}

func expiration(serial string) {
	resp, err := http.Get(url + "/time?serial=" + serial)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	btime, _ := ioutil.ReadAll(resp.Body)
	day, err := strconv.Atoi(string(btime))
	if int(day) < 0 {
		fmt.Println("License time has expired")
		os.Exit(0)
	}

	fmt.Printf("%s days left until the license expires\n", string(btime))
}
