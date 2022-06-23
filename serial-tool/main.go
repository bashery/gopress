package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(newSerial())
}

func newSerial() (serial string) {
	chars := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "l", "k", "j"}
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 10; i++ {
		serial += chars[rand.Intn(len(chars)-1)]
	}
	return serial
}
